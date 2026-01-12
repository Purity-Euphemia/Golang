import "strings"

func countWordsLongerThan(sentence string, n int) int {
    words := strings.Fields(sentence)
    count := 0
    for _, w := range words {
        if len(w) > n {
            count++
        }
    }
    return count
}
