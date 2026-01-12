import "strconv"

func priceToString(price float64) string {
    return strconv.FormatFloat(price, 'f', 2, 64)
}
