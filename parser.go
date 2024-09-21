package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// func parseHTML(html string) []string {
// 	var vulnerabilities []string

// 	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	doc.Find(".vulnerability").Each(func(i int, s *goquery.Selection) {
// 		severity := s.Find(".severity").Text()
// 		if severity == "Critical" || severity == "High" {
// 			vuln := s.Text()
// 			vulnerabilities = append(vulnerabilities, vuln)
// 		}
// 	})

//		return vulnerabilities
//	}
func parseHTML(html string) []string {
	var vulnerabilities []string

	// Parse the HTML content
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(html))
	if err != nil {
		log.Println("Error parsing HTML:", err)
		return vulnerabilities // Return empty list on error
	}

	// Find elements with class 'vulnerability'
	doc.Find(".vulnerability").Each(func(i int, s *goquery.Selection) {
		severity := s.Find(".severity").Text() // Find severity text

		// Check if severity is 'Critical' or 'High'
		if severity == "Critical" || severity == "High" {
			vuln := s.Find(".description").Text()           // Assuming the vulnerability description is in .description
			vulnerabilities = append(vulnerabilities, vuln) // Add to list of vulnerabilities
		}
	})
	fmt.Println(vulnerabilities)
	return vulnerabilities
}
