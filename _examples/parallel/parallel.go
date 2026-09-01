package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector(

		colly.MaxDepth(2),
		colly.Async(),
	)

	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 2})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		fmt.Println(link)

		e.Request.Visit(link)
	})

	c.Visit("https://en.wikipedia.org/")

	c.Wait()
}
