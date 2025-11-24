package piscine

func IterativePower(nb int, power int) int {
	if nb <= 0 {
		return 0
}
	result := 1	
	for count := 1; count <= power; count++ {
		
	result = result * nb
}

	return result

}