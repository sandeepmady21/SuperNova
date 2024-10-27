package main

import (
	"fmt"
)

func main() {
	str := "Hello, 世界, سلام دنیا"
	runes := []rune(str) // Convert string to rune slice to access characters

	for i, r := range runes {
		fmt.Printf("Rune %d: %c (Unicode: %U)\n", i, r, r)
	}
}
