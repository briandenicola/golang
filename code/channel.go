package main

import "fmt"

func main() {
    messages := make(chan string)
    go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
	
	msgChan := make(chan string)
	go func() {
		for {
			msg := <-msgChan
			fmt.Println(msg)
			if msg == "break" {
				break
			}
		}
	}()

	msgChan <- "Hello"
	msgChan <- "Good-Bye"
	msgChan <- "break"
	//msgChan <- "Test" -- Will not get called. Errors with "all goroutines are asleep. deadlock"
}