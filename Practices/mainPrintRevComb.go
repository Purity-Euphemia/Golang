package main

import "fmt"

func main() {
	first := true

	for a := 9; a >= 0; a-- {
		for b := 9; b >= 0; b-- {
			for c := 9; c >= 0; c-- {
				if a > b && b > c {

					if !first {
						fmt.Print(", ")
					}
					first = false

					fmt.Printf("%d%d%d", a, b, c)
				}
			}
		}
	}

	fmt.Println()
}
