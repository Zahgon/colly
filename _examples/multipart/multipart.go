package main

import (
	"fmt"
	"time"

	"github.com/gocolly/colly/v2"
)

func generateFormData() map[string][]byte { _ = "STUB: not implemented"; return nil }

func setupServer() { _ = "STUB: not implemented"; return }

func main() {

	setupServer()

	c := colly.NewCollector(colly.AllowURLRevisit(), colly.MaxDepth(5))

	c.OnHTML("html", func(e *colly.HTMLElement) {
		fmt.Println(e.Text)
		time.Sleep(1 * time.Second)
		e.Request.PostMultipart("http://localhost:8080/", generateFormData())
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Posting gocolly.jpg to", r.URL.String())
	})

	c.PostMultipart("http://localhost:8080/", generateFormData())
	c.Wait()
}
