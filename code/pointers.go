package main 

import "fmt"
func f( z *int ) {
	*z = 0
}

func main() {
	x := 5
	f(&x)
	fmt.Println(x)
}