package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector()

	c.OnHTML("*", func(e *colly.HTMLElement) {
		fmt.Println(e)
	})

	c.OnError(func(r *colly.Response, err error) {
		fmt.Println("Request URL:", r.Request.URL, "failed with response:", r, "\nError:", err)
	})

	c.Visit("https://definitely-not-a.website/")
}
