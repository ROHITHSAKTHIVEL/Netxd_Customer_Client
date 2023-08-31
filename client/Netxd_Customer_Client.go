package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/ROHITHSAKTHIVEL/Netxd_Customer_Proto/proto"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect :%v", err)
	}
	defer conn.Close()

	client := pb.NewCustomerServiceClient(conn)

	response, err := client.CreateCustomer(context.Background(), &pb.Customer{
		CustomerId: 4545,
		FirstName:  "Rohith ",
		SecondName: "S",
		BankId:     1,
		Balance:    1000,
		CreatedAt:  "2023.08.31",
		UpdatedAt:  "2023.08.31",
		IsActive:   true,
	})
	if err != nil {
		log.Fatalf("Failed to create customer :%v", err)
	}

	fmt.Printf("Response from customer : %v\n", response.CustomerId)
}
