package server

import (
	"context"
	"example1/pb"
)

// Calculator : Implements Calculator Service
type Calculator struct{}

// Multiply Implementation of Multiply interface
func (c *Calculator) Multiply(ctx context.Context, in *pb.TwoNumbers) (*pb.Number, error) {
	return &pb.Number{Num: in.Num1 * in.Num2}, nil
}
