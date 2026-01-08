import "strings"

func longestWord(sentence string) string {
    words := strings.Fields(sentence)
    longest := ""
    for _, word := range words {
        if len(word) > len(longest) {
            longest = word
        }
    }
    return longest
}
