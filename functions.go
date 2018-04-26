package main 

import "fmt"
func average( xs []float64 ) float64 {
	var total float64 = 0
	for _, value := range xs {
		total += value
	}
	return total / float64(len(xs))
}

func f() (int, int) {
	return 5, 6
}

func main() {
	x := []float64 {98, 93,100, 24324,83}
	fmt.Println(average(x))

	y,z := f()
	fmt.Println(z,y)

	//Closure #1
	add := func(x, y int) int {
		return x+y
	}
	fmt.Println(add(1,1))

	//Closure #2
	b := 0
	incr := func() int {
		b++
		return b
	}
	fmt.Println(incr())
	fmt.Println(incr())
}