import "strings"

func countWordsStartingWith(sentence, letter string) int {
    count := 0
    words := strings.Fields(sentence)
    for _, word := range words {
        if strings.HasPrefix(strings.ToLower(word), strings.ToLower(letter)) {
            count++
        }
    }
    return count
}
