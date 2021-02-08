package main

import (
	"context"
	"example1/pb"
	"fmt"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(":8081", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := pb.NewCalculatorClient(conn)

	val, err := client.Multiply(context.Background(), &pb.TwoNumbers{Num1: 3, Num2: 4})
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
}
