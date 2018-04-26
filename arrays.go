package main 

import "fmt"

func main() {
	x := []float64 {
		98,
		93,
		100,
		24324,
		83,
	}

	var total float64 = 0
	for i := 0; i < len(x); i++ {
		total += x[i]
	}
	fmt.Println( total / float64(len(x)))

	var total2 float64 = 0
	for _, value := range x {
		total2 += value
	}
	fmt.Println( total2 / float64(len(x)))
}