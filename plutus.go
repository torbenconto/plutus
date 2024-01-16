package plutus

import "github.com/gocolly/colly/v2"

var DefaultCollector = colly.NewCollector(
	colly.UserAgent("plutus"),
	colly.IgnoreRobotsTxt(),
	colly.MaxDepth(1),
	colly.CacheDir("./.plutus_cache"),
)
