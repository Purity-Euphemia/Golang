import "strconv"

func intSliceToString(nums []int) []string {
    result := []string{}
    for _, n := range nums {
        result = append(result, strconv.Itoa(n))
    }
    return result
}
