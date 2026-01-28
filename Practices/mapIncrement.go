import "strings"

func mapIncrement(m map[string]int) {
    for k := range m {
        m[k]++
    }
}
