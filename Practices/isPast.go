func isPast(t time.Time) bool {
    return t.Before(time.Now())
}
