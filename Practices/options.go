package main

import (
	"fmt"
	"os"
)

func printHelp() {
	fmt.Println("options: abcdefghijklmnopqrstuvwxyz")
}

func printBits(bits uint32) {
	for i := 3; i >= 0; i-- {
		byteVal := (bits >> (i * 8)) & 0xFF
		fmt.Printf("%08b", byteVal)
		if i != 0 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printHelp()
		return
	}

	var bits uint32

	for _, arg := range args {
		if len(arg) < 2 || arg[0] != '-' {
			fmt.Println("Invalid Option")
			return
		}

		if arg == "-h" {
			printHelp()
			return
		}

		for i := 1; i < len(arg); i++ {
			c := arg[i]

			if c < 'a' || c > 'z' {
				fmt.Println("Invalid Option")
				return
			}

			if c == 'h' && i == 1 {
				printHelp()
				return
			}

			bits |= 1 << (c - 'a')
		}
	}

	printBits(bits)
}
