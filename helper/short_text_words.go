package helper

import "strings"

func ShortTextWords(n int, text string) string {

	words := strings.Fields(text)

	var shortTextWords []string
	if len(words) > n {
		shortTextWords = words[:n]
	} else {
		shortTextWords = words
	}
	shortText := strings.Join(shortTextWords, " ")

	return shortText

}

func LimitCharacters(input string, limit int) string {
	if len(input) > limit {
		return input[:limit] + "..."
	}
	return input
}
