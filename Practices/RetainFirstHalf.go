package main

func RetainFirstHalf(str string) string {
	length := len(str)

	if str == 0 {
		return ""
} 	if str == 1 {
		return str
}
	half := length / 2
	return str[:half]

}