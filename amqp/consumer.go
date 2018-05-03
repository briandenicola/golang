package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"amqp/contracts"
	"github.com/streadway/amqp"
)

var (
	amqpURI = flag.String("amqp", "amqp://guest:guest@k8s03:30545/", "AMQP URI")
)

var conn *amqp.Connection
var ch *amqp.Channel
var replies <-chan amqp.Delivery

func main() {
	flag.Parse()
	fmt.Println("Start consuming the Queue...")
	var count int = 1

	var err error
	var q amqp.Queue

	conn, err = amqp.Dial(*amqpURI)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("got Connection, getting Channel...")

	ch, err = conn.Channel()
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("got Channel, declaring Exchange (%s)", "bjdExchange")

	err = ch.ExchangeDeclare(
		"bjdExchange", 			// name of the exchange
		"direct",           // type
		true,               // durable
		false,              // delete when complete
		false,              // internal
		false,              // noWait
		nil,                // arguments
	)

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("declared Exchange, declaring Queue (%s)", "bjdqueue")

	q, err = ch.QueueDeclare(
		"bjdqueue", // name, leave empty to generate a unique name
		true,            // durable
		false,           // delete when usused
		false,           // exclusive
		false,           // noWait
		nil,             // arguments
	)

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("declared Queue (%q %d messages, %d consumers), binding to Exchange (key %q)", q.Name, q.Messages, q.Consumers, "key")

	err = ch.QueueBind(
		q.Name,             // name of the queue
		"key",      // bindingKey
		"bjdExchange", // sourceExchange
		false,              // noWait
		nil,                // arguments
	)

	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("Queue bound to Exchange, starting Consume (consumer tag %q)", "bjdExample")

	replies, err = ch.Consume(
		q.Name,            // queue
		"bjdExample", // consumer
		true,              // auto-ack
		false,             // exclusive
		false,             // no-local
		false,             // no-wait
		nil,               // args
	)
	if err != nil {
		panic(err.Error())
	}

	for r := range replies {
		fmt.Printf("Reply %d -- ", count)
		user := contracts.User{}
		json.Unmarshal(r.Body, &user)
		fmt.Printf("FirstName: %s, LastName: %s\n", user.FirstName, user.LastName)
		count++
	}
}
