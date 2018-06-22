package main
//Example from Building RESTful Web Services with Go by Naren Yellavula

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "datafiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)


const (
	port = ":5051"
	noOfSteps = 3
)

type server struct {}

func (s *server) MakeTransaction(in *pb.TransactionRequest, stream pb.MoneyTransaction_MakeTransactionServer) error {
	log.Printf("Got request for money transfer . . .")
	log.Printf("Amount: %f, From A/c:%s, To A/c:%s", in.Amount, in.From, in.To)
	
	for i := 0; i < noOfSteps; i++ {
		time.Sleep(time.Second*2)
		if err := stream.Send( &pb.TransactionResponse{Status: "good",
			Step:  int32(i),
			Description: fmt.Sprintf("Description of step %d", int(i))}); err != nil { log.Fatalf("%vSend(%v) = %v", stream, "status", err)
		}
	}
	log.Printf("Successfully transfered amount $%v from %v to %v", in.Amount, in.From, in.To)
	return nil 
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMoneyTransactionServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

