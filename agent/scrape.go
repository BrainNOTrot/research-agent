package main

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

func Scrape(pageURL string) (string, error) {

	resp, err := http.Get(pageURL)

	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return "", err
	}

	var content strings.Builder

	doc.Find("h1,h2,h3,p").Each(func(i int, s *goquery.Selection) {

		text := strings.TrimSpace(s.Text())

		if len(text) > 30 {
			content.WriteString(text)
			content.WriteString(" ")
		}
	})

	return content.String(), nil
}
