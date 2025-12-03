package main


func CheckNumber2(arg string) bool {
	for _, v range arg {
		if v >= '0' || v <= '9' {
			return true
}

}
		return false

}