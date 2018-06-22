package main 
//Example from Building RESTful Web Services with Go by Naren Yellavula

import (
	"log"
	pb "datafiles"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address = "localhost:5051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}

	defer conn.Close();

	c := pb.NewMoneyTransactionClient(conn)
	from := "1234"
	to := "5678"
	amount := float32(100.00)

	r, err := c.MakeTransaction(context.Background(), &pb.TransactionRequest{From: from, To: to, Amount: amount})
	if err != nil {
		log.Fatalf("Failed with the transaction: %v", err)
	}
	log.Printf("Transaction confirmed: %t", r.Confirmation)
}