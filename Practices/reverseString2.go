func reverseString2(s string) string {
    r := ""
    for _, c := range s { r = string(c) + r }
    return r
}
