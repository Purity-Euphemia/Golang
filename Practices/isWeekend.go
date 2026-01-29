import "fmt"

func isWeekend(day time.Weekday) bool {
    return day == time.Saturday || day == time.Sunday
}
