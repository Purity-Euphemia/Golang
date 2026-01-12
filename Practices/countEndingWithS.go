import "strings"

func countEndingWithS(sentence string) int {
    words := strings.Fields(sentence)
    count := 0
    for _, w := range words {
        if strings.HasSuffix(strings.ToLower(w), "s") {
            count++
        }
    }
    return count
}
