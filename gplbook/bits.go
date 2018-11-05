package main

import (
	"crypto/sha256"
	"fmt"
)

var (
	s1 string = "my first string"
	s2 string = "my second string"
	h1        = sha256.Sum256([]byte(s1))
	h2        = sha256.Sum256([]byte(s2))
)

func main() {
	fmt.Printf("s1: %s h1: %X h1 type: %T\n", s1, h1, h1)
	fmt.Printf("s2: %s h2: %X h2 type: %T\n", s2, h2, h2)
	fmt.Printf("Number of different bits: %d\n", DifferentBits(h1, h2))
}

func bitCount(x byte) int {
	count := 0
	for x != 0 {
		x &= x - 1 //AND x with x-1 to walk through set bits 
		count++
	}
	return count
}

func DifferentBits(c1 [32]byte, c2 [32]byte) int {
	var counter int
	for x := range c1 {
		counter += bitCount(c1[x] ^ c2[x]) //XOR to get difference bits
	}
	return counter
}