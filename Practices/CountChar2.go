package main


func CountChar2(str string c rune) int {
	if str == "" {
		return 0
	}
	count := 0
	for _, a := range str {
		if a == c {
			count++
}
}
	return count
	
}