package plutus

import "github.com/gocolly/colly/v2"

const Version = "v1.4.4"

var DefaultCollector = colly.NewCollector(
	colly.UserAgent("plutus/"+Version),
	colly.IgnoreRobotsTxt(),
	colly.MaxDepth(1),
	colly.CacheDir("./.plutus_cache"),
)
