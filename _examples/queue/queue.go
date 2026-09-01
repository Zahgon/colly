package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
	"github.com/gocolly/colly/v2/queue"
)

func main() {
	url := "https://httpbin.org/delay/1"

	c := colly.NewCollector(colly.AllowURLRevisit())

	q, _ := queue.New(
		2,
		&queue.InMemoryQueueStorage{MaxSize: 10000},
	)

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("visiting", r.URL)
		if r.ID < 15 {
			r2, err := r.New("GET", fmt.Sprintf("%s?x=%v", url, r.ID), nil)
			if err == nil {
				q.AddRequest(r2)
			}
		}
	})

	for i := 0; i < 5; i++ {

		q.AddURL(fmt.Sprintf("%s?n=%d", url, i))
	}

	q.Run(c)

}
