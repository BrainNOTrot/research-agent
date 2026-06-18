package main

func PickNextTopic(
	keywords []string,
	rootTopic string,
	memory *Memory,
) string {

	bestKeyword := ""
	bestScore := -1

	for _, word := range keywords {

		if memory.HasVisited(word) {
			continue
		}

		score := ScoreKeyword(
			word,
			rootTopic,
		)

		if score > bestScore {
			bestScore = score
			bestKeyword = word
		}
	}

	return bestKeyword
}
