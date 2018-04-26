package main 

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("Hello World from goroutine") 
}

func main() {
	go f()
	fmt.Println("Hello World from main") 

	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World from main 2") 
}a