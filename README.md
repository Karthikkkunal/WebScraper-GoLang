# Vulnerability WebScraper-GoLang
The Vulnerability WebScraper in GoLang is a powerful project designed to scrape web data and identify potential vulnerabilities by leveraging the Go programming language's efficiency and concurrency. Using libraries like Colly, it enables structured and reliable data extraction from OEM websites or vulnerability databases, focusing on high and critical severity issues. The project demonstrates key GoLang concepts such as goroutines for parallel scraping, HTTP request customization, error handling, and robust data management with JSON and CSV formats. By combining cybersecurity insights, ethical scraping practices, and Go's modular design, this project offers a practical foundation for building scalable, efficient tools for vulnerability assessment and reporting.

# Before proceeding with this web scraping guide, ensure you have the necessary tools installed.

# Set Up the Environment

# Here are the prerequisites you have to meet for this Project:
•	Go 1.19+: Any version of Go greater than or equal to 1.19 is okay. You'll see version 1.22.0 in action in this tutorial.
•	A Go IDE: Visual Studio Code with the Go extension is recommended.
Open the links above to download, install, and set up the required tools by following their installation wizards.

# Set Up a Go Project
After installing Go, it's time to initialize your Golang web scraper project. Create a web-scraper-go folder and enter it in your terminal:

# Terminal
mkdir web-scraper-go 
cd web-scraper-go

# Launch a Go module with the following command:

# Terminal
go mod init scraper
The init command will initialize a scraper Go module inside your project root folder. You should now see a go.mod file with the following content in your root folder:
go.mod
module scraper
go 1.22.0

# To the Run the project :
go run . 
