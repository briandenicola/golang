package main 

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		if i % 2 == 0 {
			fmt.Println(i, " is even")
		} else {
			fmt.Println(i, " is odd")
		}
		i += 1
	}
}