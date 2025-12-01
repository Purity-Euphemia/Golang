package piscine

func Itoa(n int) string {
    if n == 0 {
        return "0"
    }

    sign := ""
    if n < 0 {
        sign = "-"
        n = -n
    }

    result := ""

    for n > 0 {
        digit := n % 10
        char := rune(digit + '0')
        result = string(char) + result
        n /= 10
    }

    return sign + result
}
