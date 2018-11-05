package main

import (
	"fmt"
)

func rotate(s []int, i int ) {
	tmp := make([]int, i)
	copy(tmp, s[:i])
	copy(s,s[i:])
	copy(s[len(s)-i:], tmp)
}

func main() {
	a := []int{0,1,2,3,4,5}
	rotate(a,2)
	fmt.Println(a)
}