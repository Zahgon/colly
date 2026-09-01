package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/debug"
)

func main() {
	url := "https://httpbin.org/delay/2"

	c := colly.NewCollector(

		colly.Debugger(&debug.LogDebugger{}),
		colly.Async(),
	)

	c.Limit(&colly.LimitRule{
		DomainGlob:  "*httpbin.*",
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	for i := 0; i < 4; i++ {
		c.Visit(fmt.Sprintf("%s?n=%d", url, i))
	}

	c.Visit(url)

	c.Wait()
}
