import "strconv"

func stringToInt(s string) int {
    num, _ := strconv.Atoi(s)
    return num
}
