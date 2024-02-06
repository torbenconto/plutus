package main

import (
	"fmt"
	"github.com/torbenconto/plutus/news"
)

func main() {
	data, err := news.NewNews("GOOG")
	if err != nil {
		panic(err)
	}

	for _, article := range data.Articles {
		fmt.Println(article.Title)
		fmt.Println(article.Link)
		fmt.Println(article.Publisher)
		fmt.Println(article.ProviderPublishTime)
		fmt.Println(article.Type)
		fmt.Println(article.RelatedTickers)
		fmt.Println()
	}
}
