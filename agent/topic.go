package main

func PickNextTopic(
	keywords []string,
	memory *Memory,
) string {

	for _, word := range keywords {

		if !memory.HasVisited(word) {

			return word
		}
	}

	return ""
}
