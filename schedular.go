package main

// import (
// 	"time"

// 	"github.com/robfig/cron/v3"
// 	logr "github.com/sirupsen/logrus"
// )

// func startScheduler() {
// 	c := cron.New()
// 	c.AddFunc("@daily", func() {
// 		url := "http://testhtml5.vulnweb.com"
// 		logr.Info("Running scheduled scraping")

// 		data, err := scrapeOEM(url)
// 		if err != nil {
// 			logr.Error("Error during scheduled scraping: ", err)
// 		} else {
// 			// You can also write this data to a file using writeReport if needed
// 			logr.Info("Scraping result: ", data)
// 		}
// 	})
// 	c.Start()

// 	// Keep the scheduler running
// 	for {
// 		time.Sleep(1 * time.Second)
// 	}
// }
