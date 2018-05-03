package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"math/rand"
	"amqp/contracts"
	"github.com/streadway/amqp"
	"time"
)

var (
	amqpURI = flag.String("amqp", "amqp://guest:guest@k8s03:30545/", "AMQP URI")
)

var conn *amqp.Connection
var ch *amqp.Channel
const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func randomString(l int) string {
	bytes := make([]byte, l)
	for i := 0; i < l; i++ {
		bytes[i] = letters[rand.Intn(len(letters))]
	}
	return string(bytes)
}

func publishMessages(messages int) {
	for i := 0; i < messages; i++ {
		user := contracts.User{}
		user.FirstName = randomString(5)
		user.LastName = randomString(10)

		payload, err := json.Marshal(user)
		if err != nil {
			panic(err.Error())
		}

		err = ch.Publish(
			"bjdExchange", 
			"key",             
			false,              
			false,              
			amqp.Publishing{
				DeliveryMode: amqp.Transient,
				ContentType:  "application/json",
				Body:         payload,
				Timestamp:    time.Now(),
			})

		if err != nil {
			panic(err.Error())
		}
	}
}

func main() {
	flag.Parse()
	fmt.Println("Starting publisher...")

	var err error

	conn, err = amqp.Dial(*amqpURI)
	if err != nil {
		panic(err.Error())
	}

	ch, err = conn.Channel()
	if err != nil {
		panic(err.Error())
	}

	err = ch.ExchangeDeclare(
		"bjdExchange",   // name
		"direct",        // type
		true,            // durable
		false,           // auto-deleted
		false,           // internal
		false,           // noWait
		nil,             // arguments
	)
	if err != nil {
		panic(err.Error())
	}

	publishMessages(100)
	defer ch.Close()
	defer conn.Close()
}
