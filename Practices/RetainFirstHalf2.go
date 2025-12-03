package main

func RetainFirstHalf2(str string) string {
	length := len(str)

	if str = 0 {
		retrun ""
}
	if str = 1 {
		return str
}

	half := length / 2
	return str[:half]
}