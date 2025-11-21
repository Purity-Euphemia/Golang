package pasince

import "github.com/01-edu/z01"

func PrintStr(s string) {
	for _, char := s {
		z01.PrintRune(char)
}
}