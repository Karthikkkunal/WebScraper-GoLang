package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

// func scrapeOEM(url string) (string, error) {
// 	c := colly.NewCollector(
// 		colly.AllowedDomains("www.itsecgames.com"),
// 	)

// 	var result string
// 	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
// 		link := e.Attr("href")
// 		result += fmt.Sprintf("Link found: %q -> %s\n", e.Text, link)
// 	})

// 	c.OnRequest(func(r *colly.Request) {
// 		fmt.Println("Visiting", r.URL.String())
// 	})

// 	err := c.Visit(url)
// 	if err != nil {
// 		return "", err
// 	}

// 	return result, nil

// }
func scrapeOEM(url string) (string, error) {
	c := colly.NewCollector()

	var htmlContent string

	// Callback to store the full HTML content of the page
	c.OnResponse(func(r *colly.Response) {
		htmlContent = string(r.Body) // Convert response body to string
	})

	// Log requests being made (optional)
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})

	// Visit the URL
	err := c.Visit(url)
	if err != nil {
		return "", err
	}

	// Return the scraped HTML content
	return htmlContent, nil
}
