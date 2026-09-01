package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gocolly/colly/v2"
)

type item struct {
	StoryURL  string
	Source    string
	comments  string
	CrawledAt time.Time
	Comments  string
	Title     string
}

func main() {
	stories := []item{}

	c := colly.NewCollector(

		colly.AllowedDomains("old.reddit.com"),

		colly.Async(true),
	)

	c.OnHTML(".top-matter", func(e *colly.HTMLElement) {
		temp := item{}
		temp.StoryURL = e.ChildAttr("a[data-event-action=title]", "href")
		temp.Source = "https://old.reddit.com/r/programming/"
		temp.Title = e.ChildText("a[data-event-action=title]")
		temp.Comments = e.ChildAttr("a[data-event-action=comments]", "href")
		temp.CrawledAt = time.Now()
		stories = append(stories, temp)
	})

	c.OnHTML("span.next-button", func(h *colly.HTMLElement) {
		t := h.ChildAttr("a", "href")
		c.Visit(t)
	})

	c.Limit(&colly.LimitRule{
		Parallelism: 2,
		RandomDelay: 5 * time.Second,
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())

	})

	reddits := os.Args[1:]
	for _, reddit := range reddits {
		c.Visit(reddit)

	}

	c.Wait()
	fmt.Println(stories)

}
