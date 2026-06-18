package main

import "strings"

func ScoreKeyword(
	keyword string,
	rootTopic string,
) int {

	score := 0

	rootWords := strings.Fields(
		strings.ToLower(rootTopic),
	)

	keyword = strings.ToLower(keyword)

	for _, word := range rootWords {
		if strings.Contains(keyword, word) {
			score += 10
		}
	}

	return score
}
