func gradeScore(score int) string {
    if score >= 70 {
        return "A"
    } else if score >= 60 {
        return "B"
    } else if score >= 50 {
        return "C"
    }
    return "F"
}
