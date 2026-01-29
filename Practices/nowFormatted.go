import "time"

func nowFormatted() string {
    return time.Now().Format("2006-01-02")
}
