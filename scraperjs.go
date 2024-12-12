package main

import (
	"context"

	"github.com/chromedp/chromedp"
)

func scrapeWithChromedp(url string) (string, error) {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var result string
	err := chromedp.Run(ctx,
		chromedp.Navigate(url),
		chromedp.WaitVisible(`body`, chromedp.ByQuery),
		chromedp.OuterHTML(`html`, &result),
	)
	if err != nil {
		return "", err
	}

	return result, nil
}
