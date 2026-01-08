func secondsToMinSec(seconds int) (int, int) {
    minutes := seconds / 60
    remaining := seconds % 60
    return minutes, remaining
}
