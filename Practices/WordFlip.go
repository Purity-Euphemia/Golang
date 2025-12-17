package piscine

import "strings"

func WordFlip(str string) string {
	trimmed := strings.TrimSpace(str)

	if trimmed == "" {
		return "Invalid Output\n"
	}

	words := strings.Fields(trimmed)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}

	return strings.Join(words, " ") + "\n"
}
