import "strconv"

func toBinary(n int) string {
    return strconv.FormatInt(int64(n), 2)
}
