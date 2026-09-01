package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

func main() {
	url := "https://httpbin.org/delay/2"

	c := colly.NewCollector(

		colly.Async(),

		colly.Debugger(&debug.LogDebugger{}),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
	})

	for i := 0; i < 5; i++ {
		c.Visit(fmt.Sprintf("%s?n=%d", url, i))
	}

	c.Wait()
}
