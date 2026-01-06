func convertMinutes(minutes int) (int, int) {
    hours := minutes / 60
    remaining := minutes % 60
    return hours, remaining
}
