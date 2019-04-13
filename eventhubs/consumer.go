package main

import (
	"context"
	"fmt"
	"os"
	"time"
	"github.com/Azure/azure-event-hubs-go"
)

const (
	HubName           = "producerConsumer"
)

func main() {
	hub := initHub()
	exit := make(chan struct{})

	handler := func(ctx context.Context, event *eventhub.Event) error {
		text := string(event.Data)
		if text == "exit\n" {
			fmt.Println("Oh snap!! Someone told me to exit!")
			exit <- *new(struct{})
		} else {
			fmt.Println(string(event.Data))
		}
		return nil
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	runtimeInfo, err := hub.GetRuntimeInformation(ctx)
	if err != nil {
		panic(err)
	}
	for _, partitionID := range runtimeInfo.PartitionIDs {
		hub.Receive(ctx, partitionID, handler, eventhub.ReceiveWithLatestOffset())
	}
	cancel()

	fmt.Println("I am listening...")

	select {
	case <-exit:
		fmt.Println("closing after 2 seconds")
		select {
		case <-time.After(2 * time.Second):
			return
		}
	}
}

func initHub() (*eventhub.Hub) {
	namespace := mustGetenv("EVENTHUB_NAMESPACE")
        hub,err := eventhub.NewHubFromConnectionString(namespace)
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
