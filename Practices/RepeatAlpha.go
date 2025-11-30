package piscine

func RepeatAlpha(s string) string {
    result := ""

    for _, r := range s {
        if r >= 'a' && r <= 'z' {
            repeat := int(r - 'a' + 1)
            for i := 0; i < repeat; i++ {
                result += string(r)
            }
        } else if r >= 'A' && r <= 'Z' {
            repeat := int(r - 'A' + 1)
            for i := 0; i < repeat; i++ {
                result += string(r)
            }
        } else {
		result += string(r)
        }
    }

    return result
}
