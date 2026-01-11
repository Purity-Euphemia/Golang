import "strings"

func countLongWords(sentence string) int {
    words := strings.Fields(sentence)
    count := 0
    for _, w := range words {
        if len(w) > 5 {
            count++
        }
    }
    return count
}
