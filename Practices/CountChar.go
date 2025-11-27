package main

func CountChar(str string, c rune) int {

	if str <= "" {
		return 0

}
	count := 0

	for _, ch := range str {
		if ch == c {
			count++

}

}
	return count
}