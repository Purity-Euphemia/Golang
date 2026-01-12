import "strings"

func countWord(sentence, word string) int {
    count := 0
    for _, w := range strings.Fields(sentence) {
        if w == word {
            count++
        }
    }
    return count
}
