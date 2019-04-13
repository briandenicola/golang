package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/Azure/azure-event-hubs-go"
)

const (
	HubName           = "producerConsumer"
)

func main() {
	hub := initHub()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		err := hub.Send(ctx, eventhub.NewEventFromString(text))
		if err != nil {
			log.Fatal(err)
		}
		if text == "exit\n" {
			break
		}
		cancel()
	}
}

func initHub() (*eventhub.Hub) {
	namespace := mustGetenv("EVENTHUB_NAMESPACE")
	hub, err := eventhub.NewHubFromConnectionString(namespace)
	if err != nil { 
		panic(err)
	}
	return hub
}

func mustGetenv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic("Environment variable '" + key + "' required for integration tests.")
	}
	return v
}
