import "strings"

func countLines(text string) int {
    return len(strings.Split(text, "\n"))
}
