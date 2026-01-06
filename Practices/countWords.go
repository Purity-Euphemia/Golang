import "strings"

func countWords(sentence string) int {
    words := strings.Fields(sentence)
    return len(words)
}
