import "math"

func isPerfectSquare(n int) bool {
    root := int(math.Sqrt(float64(n)))
    return root*root == n
}
