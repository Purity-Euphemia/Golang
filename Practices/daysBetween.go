import "time"

func daysBetween(a, b time.Time) int {
    diff := b.Sub(a)
    return int(diff.Hours() / 24)
}
