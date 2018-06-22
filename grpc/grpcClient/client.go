package main 
//Example from Building RESTful Web Services with Go by Naren Yellavula

import (
	"io"
	"log"

	pb "datafiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5051"
)

func ReceiveStream(client pb.MoneyTransactionClient, request *pb.TransactionRequest) {
	log.Println("Started listening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTransactio(_) = _, %v", client, err)
	}
	
	for {
		response, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}
		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}
}

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close();
	client := pb.NewMoneyTransactionClient(conn)

	from := "1234"
	to := "5678"
	amount := float32(100.00)

	ReceiveStream(client, &pb.TransactionRequest{From: from, To: to, Amount: amount})
}