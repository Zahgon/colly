package main

import (
	"encoding/json"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gocolly/colly/v2"
)

type Course struct {
	Title       string
	Description string
	Creator     string
	Level       string
	URL         string
	Language    string
	Commitment  string
	Rating      string
}

func main() {
	fName := "courses.json"
	file, err := os.Create(fName)
	if err != nil {
		log.Fatalf("Cannot create file %q: %s\n", fName, err)
		return
	}
	defer file.Close()

	c := colly.NewCollector(

		colly.AllowedDomains("coursera.org", "www.coursera.org"),

		colly.CacheDir("./coursera_cache"),

		colly.CacheExpiration(24*time.Hour),
	)

	detailCollector := c.Clone()

	courses := make([]Course, 0, 200)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		if e.Attr("class") == "Button_1qxkboh-o_O-primary_cv02ee-o_O-md_28awn8-o_O-primaryLink_109aggg" {
			return
		}
		link := e.Attr("href")

		if !strings.HasPrefix(link, "/browse") || strings.Index(link, "=signup") > -1 || strings.Index(link, "=login") > -1 {
			return
		}

		e.Request.Visit(link)
	})

	c.OnRequest(func(r *colly.Request) {
		log.Println("visiting", r.URL.String())
	})

	c.OnHTML(`a.collection-product-card`, func(e *colly.HTMLElement) {

		courseURL := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(courseURL, "coursera.org/learn") != -1 {
			detailCollector.Visit(courseURL)
		}
	})

	detailCollector.OnHTML(`div[id=rendered-content]`, func(e *colly.HTMLElement) {
		log.Println("Course found", e.Request.URL)
		title := e.ChildText(".banner-title")
		if title == "" {
			log.Println("No title found", e.Request.URL)
		}
		course := Course{
			Title:       title,
			URL:         e.Request.URL.String(),
			Description: e.ChildText("div.content"),
			Creator:     e.ChildText("li.banner-instructor-info > a > div > div > span"),
			Rating:      e.ChildText("span.number-rating"),
		}

		e.ForEach(".AboutCourse .ProductGlance > div", func(_ int, el *colly.HTMLElement) {
			svgTitle := strings.Split(el.ChildText("div:nth-child(1) svg title"), " ")
			lastWord := svgTitle[len(svgTitle)-1]
			switch lastWord {

			case "languages":
				course.Language = el.ChildText("div:nth-child(2) > div:nth-child(1)")

			case "Level":
				course.Level = el.ChildText("div:nth-child(2) > div:nth-child(1)")

			case "complete":
				course.Commitment = el.ChildText("div:nth-child(2) > div:nth-child(1)")
			}
		})
		courses = append(courses, course)
	})

	c.Visit("https://coursera.org/browse")

	enc := json.NewEncoder(file)
	enc.SetIndent("", "  ")

	enc.Encode(courses)
}
