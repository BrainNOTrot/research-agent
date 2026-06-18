package main

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

func Search(query string) ([]string, error) {

	searchURL :=
		"https://html.duckduckgo.com/html/?q=" +
			url.QueryEscape(query)

	req, _ := http.NewRequest("GET", searchURL, nil)

	req.Header.Set(
		"User-Agent",
		"Mozilla/5.0",
	)

	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)

	if err != nil {
		return nil, err
	}

	var results []string

	doc.Find(".result__a").Each(func(i int, s *goquery.Selection) {

		href, exists := s.Attr("href")

		if !exists {
			return
		}

		parsedURL, err := url.Parse(href)

		if err == nil {

			realURL := parsedURL.Query().Get("uddg")

			if realURL != "" {
				results = append(results, realURL)
			}
		}
	})

	if len(results) == 0 {
		return nil, fmt.Errorf("no results found")
	}

	if len(results) > 5 {
		results = results[:5]
	}

	return results, nil
}
