package main

import (
	"regexp"
	"sort"
	"strings"
)

var stopWords = map[string]bool{
	"the":  true,
	"and":  true,
	"for":  true,
	"with": true,
	"that": true,
	"this": true,
	"from": true,
	"have": true,
	"were": true,
	"been": true,
	"into": true,
}

func ExtractKeywords(text string) []string {

	re := regexp.MustCompile(`[^a-zA-Z ]`)
	text = re.ReplaceAllString(text, " ")

	text = strings.ToLower(text)

	words := strings.Fields(text)

	counts := map[string]int{}

	for _, word := range words {

		if len(word) < 5 {
			continue
		}

		if stopWords[word] {
			continue
		}

		counts[word]++
	}

	type pair struct {
		word  string
		count int
	}

	var pairs []pair

	for w, c := range counts {
		pairs = append(pairs, pair{w, c})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	var result []string

	for i := 0; i < len(pairs) && i < 20; i++ {
		result = append(result, pairs[i].word)
	}

	return result
}
