package helper

import (
	"regexp"
	"strconv"
	"strings"
)

const regex = `<.*?>`

func StringWithoutSpaces(inputString string) string {
	var stringWithoutSpaces string
	for _, char := range inputString {
		if char != ' ' {
			stringWithoutSpaces += string(char)
		}
	}
	return stringWithoutSpaces
}

func StrinParameterJudulID(judul string) uint {

	parts := strings.Split(judul, "-")

	if len(parts) > 0 {
		stringInt := parts[len(parts)-1]
		number, err := strconv.ParseUint(stringInt, 10, 0)
		if err != nil {
			return 0
		}

		return uint(number)
	} else {
		return 0
	}

}

func StripHtmlRegex(s string) string {
	r := regexp.MustCompile(regex).ReplaceAllString(s, "")
	return strings.NewReplacer(">", "", "<", "").Replace(r)
}
