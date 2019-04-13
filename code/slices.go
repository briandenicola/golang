package main 

import "fmt"

func main() {
	x := []int {
		98,
		93,
		100,
		24324,
		83}

	y1 := x[1:3]
	for z := range y1 {
		fmt.Printf("y1[%d]: %d\n", z, y1[z])
	}

	y2 := x[:3]
	for z := range y2 {
		fmt.Printf("y2[%d]: %d\n", z, y2[z])
	}

	y3 := x[3:]
	for z := range y3 {
		fmt.Printf("y3[%d]: %d\n", z, y3[z])
	}

	y4 := x[3:]
	x[3] = 15
	for z := range y4 {
		fmt.Printf("y4[%d]: %d\n", z, y4[z]) //Update x but y[3] will also print 15
	}
}