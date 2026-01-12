import "strconv"

func sumAsString(nums []int) string {
    sum := 0
    for _, n := range nums {
        sum += n
    }
    return strconv.Itoa(sum)
}
