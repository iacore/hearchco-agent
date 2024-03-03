package router

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/goccy/go-json"

	"github.com/hearchco/hearchco/src/cache"
	"github.com/hearchco/hearchco/src/config"
	"github.com/hearchco/hearchco/src/search"
	"github.com/hearchco/hearchco/src/search/category"
	"github.com/hearchco/hearchco/src/search/engines"
	"github.com/hearchco/hearchco/src/search/result"
)

func Search(c *gin.Context, db cache.DB, conf config.Config) error {
	var query, pages, deepSearch, locale, categ, useragent, safesearch, mobile string
	var ccateg category.Name

	switch c.Request.Method {
	case "", "GET":
		{
			query = c.Query("q")
			pages = c.DefaultQuery("pages", "1")
			deepSearch = c.DefaultQuery("deep", "false")
			locale = c.DefaultQuery("locale", config.DefaultLocale)
			categ = c.DefaultQuery("category", "")
			useragent = c.DefaultQuery("useragent", "")
			safesearch = c.DefaultQuery("safesearch", "false")
			mobile = c.DefaultQuery("mobile", "false")
		}
	case "POST":
		{
			query = c.PostForm("q")
			pages = c.DefaultPostForm("pages", "1")
			deepSearch = c.DefaultPostForm("deep", "false")
			locale = c.DefaultPostForm("locale", config.DefaultLocale)
			categ = c.DefaultPostForm("category", "")
			useragent = c.DefaultPostForm("useragent", "")
			safesearch = c.DefaultPostForm("safesearch", "false")
			mobile = c.DefaultPostForm("mobile", "false")
		}
	}

	if query == "" {
		c.String(http.StatusOK, "")
	} else {
		maxPages, pageserr := strconv.Atoi(pages)
		if pageserr != nil {
			c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Cannot convert pages value (%q) to int", pages))
			return fmt.Errorf("router.Search(): cannot convert pages value %q to int: %w", pages, pageserr)
		}

		visitPages, deeperr := strconv.ParseBool(deepSearch)
		if deeperr != nil {
			c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Cannot convert deep value (%q) to bool", deepSearch))
			return fmt.Errorf("router.Search(): cannot convert deep value %q to int: %w", deepSearch, deeperr)
		}

		if lerr := engines.ValidateLocale(locale); lerr != nil {
			c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Invalid locale value (%q), should be of the form \"en_US\"", locale))
			return fmt.Errorf("router.Search(): invalid locale value %q: %w", locale, lerr)
		}

		ccateg = category.SafeFromString(categ)
		if ccateg == category.UNDEFINED {
			c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Invalid category value (%q)", categ))
			return fmt.Errorf("router.Search(): invalid category value %q", categ)
		}

		safeSearchB, safeerr := strconv.ParseBool(safesearch)
		if safeerr != nil {
			c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Cannot convert safesearch value (%q) to bool", safesearch))
			return fmt.Errorf("router.Search(): cannot convert safesearch value %q to bool: %w", safesearch, safeerr)
		}

		isMobile, mobileerr := strconv.ParseBool(mobile)
		if mobileerr != nil {
			c.String(http.StatusUnprocessableEntity, fmt.Sprintf("Cannot convert mobile value (%q) to bool", mobile))
			return fmt.Errorf("router.Search(): cannot convert mobile value %q to bool: %w", mobile, mobileerr)
		}

		options := engines.Options{
			MaxPages:   maxPages,
			VisitPages: visitPages,
			Category:   ccateg,
			UserAgent:  useragent,
			Locale:     locale,
			SafeSearch: safeSearchB,
			Mobile:     isMobile,
		}

		results, foundInDB := search.Search(query, options, db, conf.Settings, conf.Categories)

		resultsShort := result.Shorten(results)
		if resultsJson, err := json.Marshal(resultsShort); err != nil {
			c.String(http.StatusInternalServerError, "")
			return fmt.Errorf("router.Search(): failed marshalling results: %v\n with error: %w", resultsShort, err)
		} else {
			c.String(http.StatusOK, string(resultsJson))
		}

		search.CacheAndUpdateResults(query, options, db, conf.Server.Cache.TTL, conf.Settings, conf.Categories, results, foundInDB)
	}
	return nil
}
