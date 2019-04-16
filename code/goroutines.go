package main 

import (
	"fmt"
	"time"
)

func f() {
	fmt.Println("Hello World from goroutine") 
}

func worker(msg string ) {
    fmt.Print("working...")
    time.Sleep(time.Second)
    fmt.Println(msg)
}

func main() {
	go f()
	fmt.Println("Hello World from main") 

	go f()
	time.Sleep(1 * time.Second)
	fmt.Println("Hello World from main 2") 

	msg := "Hello World"
	go worker(msg)

	func (msg string) {
		fmt.Println(msg)
	}("Good-bye")
	
	fmt.Scanln()
}