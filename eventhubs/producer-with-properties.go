package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"
	"crypto/rand"

	"github.com/Azure/azure-event-hubs-go/v3"
)

const (
	HubName           = "traces"
)

func createUUID() (string) {
	buf := make([]byte, 16)

	if _, err := rand.Read(buf); err != nil {
		panic(err.Error())
	}

	return fmt.Sprintf("%x-%x-%x-%x-%x", buf[0:4], buf[4:6], buf[6:8], buf[8:10], buf[10:])
}

func main() {
	hub := initHub()

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		text, _ := reader.ReadString('\n')
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		evt := eventhub.NewEventFromString(text) 
		evt.Set("parentSpanId", createUUID())

		err := hub.Send(ctx,evt)
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
