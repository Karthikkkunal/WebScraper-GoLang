package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

var vulnerabilities []string

func home(w http.ResponseWriter, r *http.Request) {
	msg := "Hello, welcome to your app. Use the following suffixes on the URL to show different results:\n1) '/scrape' to show results of web scraping.\n2) '/crawl' to show results of a web crawler"
	log.Println("Received request for the home page")
	w.Write([]byte(msg))
}

func scrapeHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://testhtml5.vulnweb.com"
	log.Println("Received request to scrape URL: ", url)

	// Scrape the website
	htmlContent, err := scrapeOEM(url)
	if err != nil {
		http.Error(w, "Error occurred while scraping: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Parse the HTML to find vulnerabilities
	vulnerabilities := parseHTML(htmlContent)

	// Generate HTML content for the response
	var responseHTML string
	if len(vulnerabilities) == 0 {
		responseHTML = "<html><body><h1>No critical or high vulnerabilities found.</h1></body></html>"
	} else {
		responseHTML = "<html><body><h1>Found Vulnerabilities:</h1><ul>"
		for _, vuln := range vulnerabilities {
			responseHTML += "<li>" + vuln + "</li>"
		}
		responseHTML += "</ul></body></html>"
	}

	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(responseHTML))
}

func scrapeWithChromedpHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://testhtml5.vulnweb.com"
	log.Println("Received request to scrape with Chromedp: ", url)

	// Get HTML content using Chromedp
	htmlContent, err := scrapeWithChromedp(url)
	if err != nil {
		http.Error(w, "Error occurred while scraping with Chromedp: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve the HTML content as a new webpage
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(htmlContent))
}

func reportHandler(w http.ResponseWriter, r *http.Request) {
	url := "http://testhtml5.vulnweb.com"
	filename := "vulnerabilities_report.json"
	testVulnerabilities(url)
	fmt.Println("Vunera : ")
	fmt.Println(vulnerabilities)
	// Scrape and parse data
	htmlContent, err := scrapeOEM(url)
	if err != nil {
		http.Error(w, "Error occurred while scraping: "+err.Error(), http.StatusInternalServerError)
		return
	}
	vun := parseHTML(htmlContent)

	// Append each vulnerability to the vulnerabilities slice
	for _, v := range vun {
		vulnerabilities = append(vulnerabilities, v)
	} // Generate JSON report
	err = generateReportJSON(vulnerabilities, filename)
	if err != nil {
		http.Error(w, "Error occurred while generating report: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Serve the JSON report
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		http.Error(w, "Error occurred while reading report file: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(fileContent)
}

func testVulnerabilities(baseURL string) {
	// Test for XSS
	testXSS(baseURL)

	// Test for SQL Injection
	testSQLInjection(baseURL)

	// Test for CSRF
	testCSRF(baseURL)

	// Test for IDOR
	testIDOR(baseURL)

	// Test for Remote Code Execution (RCE)
	testRCE(baseURL)

	// Test for File Inclusion (LFI/RFI)
	testFileInclusion(baseURL)

	// Test for Security Misconfiguration
	testSecurityMisconfiguration(baseURL)

	// Test for Broken Authentication
	testBrokenAuthentication(baseURL)

	// Test for Sensitive Data Exposure
	testSensitiveDataExposure(baseURL)
}

func testXSS(urlStr string) {
	payload := "<script>alert('XSS detected!');</script>"

	// Encode the payload as a query parameter
	encodedPayload := url.QueryEscape(payload)

	// Construct the final URL with encoded query parameters
	reqURL := fmt.Sprintf("%s?query=%s", urlStr, encodedPayload)

	// Create the HTTP request
	client := &http.Client{}
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Send the HTTP request
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Check response status
	if resp.StatusCode == 200 {
		vulnerabilityMessage := fmt.Sprintf("XSS: Possible/Vulnerable; Payload: %s; Status Code: %d", payload, resp.StatusCode)
		fmt.Println(vulnerabilityMessage)
		vulnerabilities = append(vulnerabilities, vulnerabilityMessage)
	} else {
		vulnerabilityMessage := fmt.Sprintf("XSS: Not Detected; Payload: %s; Status Code: %d", payload, resp.StatusCode)
		fmt.Println(vulnerabilityMessage)
	}

}

func testSQLInjection(urlStr string) {
	payload := "' OR '1'='1"
	testWithPayload(urlStr, payload, "SQL INJECTION")
}

func testCSRF(urlStr string) {
	// Implement CSRF token testing here
}

func testIDOR(urlStr string) {
	payload := "2" // Example for IDOR testing
	testWithPayload(fmt.Sprintf("%s/resource/1", urlStr), payload, "IDOR")
}

func testRCE(urlStr string) {
	// Implement RCE testing here, e.g., by testing file upload features
}

func testFileInclusion(urlStr string) {
	payload := "../../../../etc/passwd"
	testWithPayload(urlStr, payload, "FileInclusion")
}

func testSecurityMisconfiguration(urlStr string) {
	payload := "/admin" // Example check
	testWithPayload(urlStr, payload, "SecurityMisconfiguration")
}

func testBrokenAuthentication(urlStr string) {
	// Example list of common passwords
	passwords := []string{"password", "123456", "admin", "letmein", "password1"}
	attackName := "Broken Authentication"

	for _, password := range passwords {
		loginURL := fmt.Sprintf("%s/login", urlStr)                                     // Adjust URL path as needed
		payload := fmt.Sprintf("username=admin&password=%s", url.QueryEscape(password)) // Adjust payload format as needed

		client := &http.Client{}
		req, err := http.NewRequest("POST", loginURL, strings.NewReader(payload))
		if err != nil {
			log.Println("Error creating request:", err)
			continue
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

		resp, err := client.Do(req)
		if err != nil {
			log.Println("Error sending request:", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode == 200 {
			vulnerabilityMessage := fmt.Sprintf("%s: Possible/Vulnerable; Password: %s; Status Code: %d", attackName, password, resp.StatusCode)
			fmt.Println(vulnerabilityMessage)
			vulnerabilities = append(vulnerabilities, vulnerabilityMessage)
		} else {
			fmt.Printf("%s: Not Detected; Password: %s; Status Code: %d\n", attackName, password, resp.StatusCode)
		}
	}
}

func testSensitiveDataExposure(urlStr string) {
	// Example list of keywords to check for in the response
	keywords := []string{"password", "secret", "api_key", "token", "credentials"}
	attackName := "Sensitive Data Exposure"

	client := &http.Client{}
	req, err := http.NewRequest("GET", urlStr, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	// Read response body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error reading response body:", err)
		return
	}

	// Check for sensitive data in response body
	for _, keyword := range keywords {
		if strings.Contains(string(body), keyword) {
			vulnerabilityMessage := fmt.Sprintf("%s: Possible/Vulnerable; Keyword: %s; Status Code: %d", attackName, keyword, resp.StatusCode)
			fmt.Println(vulnerabilityMessage)
			vulnerabilities = append(vulnerabilities, vulnerabilityMessage)
		}
	}
}

func testWithPayload(baseURL, payload, attackName string) {
	encodedPayload := url.QueryEscape(payload)
	reqURL := fmt.Sprintf("%s?query=%s", baseURL, encodedPayload)

	client := &http.Client{}
	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		log.Println("Error creating request:", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending request:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode == 200 {
		vulnerabilityMessage := fmt.Sprintf("%s: Possible/Vulnerable; Payload: %s; Status Code: %d", attackName, payload, resp.StatusCode)
		fmt.Println(vulnerabilityMessage)
		vulnerabilities = append(vulnerabilities, vulnerabilityMessage)
	} else {
		vulnerabilityMessage := fmt.Sprintf("%s: Not Detected; Payload: %s; Status Code: %d", attackName, payload, resp.StatusCode)
		fmt.Println(vulnerabilityMessage)
	}
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/scrape", scrapeHandler)
	http.HandleFunc("/scrapejs", scrapeWithChromedpHandler) // New handler for Chromedp
	http.HandleFunc("/report", reportHandler)

	log.Println("Server is listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
