package main

import (
	"context"
	"fmt"
	"log"

	pb "Netxd_Project/Netxd_Customer/customer"

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
		CustomerId: 123,
		FirstName:  "kishore",
		SecondName: "",
		BankId:     0,
		Balance:    10000,
		CreatedAt:  "2023.08.30",
		UpdatedAt:  "2023.08.30",
		IsActive:   true,
	})
	if err != nil {
		log.Fatalf("Failed to create customer :%v", err)
	}

	fmt.Printf("Response from customer : %v\n", response.CustomerId)
}
