package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {
	c := colly.NewCollector()

	c.OnHTML("fin-streamer", func(h *colly.HTMLElement) {

		if h.Attr("data-field") == "regularMarketPrice" && h.Attr("active") == "" {
			fmt.Println(h.Text)
		}
	})

	c.Visit("https://finance.yahoo.com/quote/amd")

}
