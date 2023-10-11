package config

import (
	"github.com/tminaorg/brzaguza/src/category"
	"github.com/tminaorg/brzaguza/src/engines"
)

const DefaultLocale string = "en-US"

func NewRanking() Ranking {
	return Ranking{
		REXP: 0.5,
		A:    1,
		B:    0,
		C:    1,
		D:    0,
		TRA:  1,
		TRB:  0,
		TRC:  1,
		TRD:  0,
		Engines: map[string]EngineRanking{
			engines.BING.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.BRAVE.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.DUCKDUCKGO.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.ETOOLS.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.GOOGLE.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.MOJEEK.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.PRESEARCH.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.QWANT.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.STARTPAGE.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.SWISSCOWS.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.YAHOO.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.YANDEX.ToLower(): {
				Mul:   1,
				Const: 0,
			},
			engines.YEP.ToLower(): {
				Mul:   1,
				Const: 0,
			},
		},
	}
}

func NewSettings() map[string]Settings {
	return map[string]Settings{
		engines.BING.ToLower(): {
			Shortcut: "bi",
		},
		engines.BRAVE.ToLower(): {
			Shortcut: "br",
		},
		engines.DUCKDUCKGO.ToLower(): {
			Shortcut: "ddg",
		},
		engines.ETOOLS.ToLower(): {
			Shortcut: "ets",
		},
		engines.GOOGLE.ToLower(): {
			Shortcut: "go",
		},
		engines.MOJEEK.ToLower(): {
			Shortcut: "mjk",
		},
		engines.PRESEARCH.ToLower(): {
			Shortcut: "ps",
		},
		engines.QWANT.ToLower(): {
			Shortcut: "qw",
		},
		engines.STARTPAGE.ToLower(): {
			Shortcut: "sp",
		},
		engines.SWISSCOWS.ToLower(): {
			Shortcut: "sc",
		},
		engines.YAHOO.ToLower(): {
			Shortcut: "yh",
		},
		engines.YEP.ToLower(): {
			Shortcut: "yep",
		},
	}
}

func NewAllEnabled() map[string]Engine {
	return map[string]Engine{
		engines.BING.ToLower(): {
			Enabled: true,
		},
		engines.BRAVE.ToLower(): {
			Enabled: true,
		},
		engines.DUCKDUCKGO.ToLower(): {
			Enabled: true,
		},
		engines.ETOOLS.ToLower(): {
			Enabled: true,
		},
		engines.GOOGLE.ToLower(): {
			Enabled: true,
		},
		engines.MOJEEK.ToLower(): {
			Enabled: true,
		},
		engines.PRESEARCH.ToLower(): {
			Enabled: true,
		},
		engines.QWANT.ToLower(): {
			Enabled: true,
		},
		engines.STARTPAGE.ToLower(): {
			Enabled: true,
		},
		engines.SWISSCOWS.ToLower(): {
			Enabled: true,
		},
		engines.YAHOO.ToLower(): {
			Enabled: true,
		},
		engines.YEP.ToLower(): {
			Enabled: true,
		},
	}
}

func NewInfo() map[string]Engine {
	return map[string]Engine{
		engines.BING.ToLower(): {
			Enabled: true,
		},
		engines.BRAVE.ToLower(): {
			Enabled: false,
		},
		engines.DUCKDUCKGO.ToLower(): {
			Enabled: false,
		},
		engines.ETOOLS.ToLower(): {
			Enabled: false,
		},
		engines.GOOGLE.ToLower(): {
			Enabled: true,
		},
		engines.MOJEEK.ToLower(): {
			Enabled: true,
		},
		engines.PRESEARCH.ToLower(): {
			Enabled: false,
		},
		engines.QWANT.ToLower(): {
			Enabled: false,
		},
		engines.STARTPAGE.ToLower(): {
			Enabled: false,
		},
		engines.SWISSCOWS.ToLower(): {
			Enabled: false,
		},
		engines.YAHOO.ToLower(): {
			Enabled: false,
		},
		engines.YEP.ToLower(): {
			Enabled: false,
		},
	}
}

func New() *Config {
	return &Config{
		Server: Server{
			Port:         3030,
			FrontendUrls: []string{"http://localhost:8000"},
			Cache: Cache{
				Type: "pebble",
				Redis: Redis{
					Host: "localhost",
					Port: 6379,
				},
			},
		},
		Settings: NewSettings(),
		Categories: map[category.Name]Category{
			category.GENERAL: {
				Engines: NewAllEnabled(),
				Ranking: NewRanking(),
			},
			category.INFO: {
				Engines: NewInfo(),
				Ranking: NewRanking(),
			},
			category.SCIENCE: {
				Engines: NewAllEnabled(),
				Ranking: NewRanking(),
			},
			category.NEWS: {
				Engines: NewAllEnabled(),
				Ranking: NewRanking(),
			},
			category.BLOG: {
				Engines: NewAllEnabled(),
				Ranking: NewRanking(),
			},
			category.SURF: {
				Engines: NewAllEnabled(),
				Ranking: NewRanking(),
			},
			category.NEWNEWS: {
				Engines: NewAllEnabled(),
				Ranking: NewRanking(),
			},
		},
	}
}
