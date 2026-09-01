package main

import (
	"fmt"

	"github.com/gocolly/colly/v2"
)

func main() {

	knownUrls := []string{}

	c := colly.NewCollector(colly.AllowedDomains("www.shopify.com"))

	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {
		knownUrls = append(knownUrls, e.Text)
	})

	c.Visit("https://www.shopify.com/sitemap.xml")

	fmt.Println("All known URLs:")
	for _, url := range knownUrls {
		fmt.Println("\t", url)
	}
	fmt.Println("Collected", len(knownUrls), "URLs")
}
