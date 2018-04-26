package main 

import (
	"fmt"
)

func main() {
	salary := make(map[string]int)
	salary["brian"] = 5000
	salary["joe"] = 32000
	fmt.Println("salary map contents:", salary)
}