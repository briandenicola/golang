package main

import "fmt"

func main() {
	var runes []rune
	for _, r := range "Hello" 	{
		runes = append(runes, r)
	}
	fmt.Printf("%q\n", runes)
}
