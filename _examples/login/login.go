package main

import (
	"log"

	"github.com/gocolly/colly/v2"
)

func main() {

	c := colly.NewCollector()

	err := c.Post("http://example.com/login", map[string]string{"username": "admin", "password": "admin"})
	if err != nil {
		log.Fatal(err)
	}

	c.OnResponse(func(r *colly.Response) {
		log.Println("response received", r.StatusCode)
	})

	c.Visit("https://example.com/")
}
