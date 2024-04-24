package main

import (
	"context"
	"fmt"

	pb "github.com/fxivan/go-grpc/proto"
	"github.com/google/uuid"
	"google.golang.org/grpc"
)

func generateID() string {
	id := uuid.New()
	return id.String()
}

func main() {
	conn, err := grpc.Dial("localhost:5051", grpc.WithInsecure())
	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	serviceClient := pb.NewWishListServiceClient(conn)

	res, err := serviceClient.Create(context.Background(), &pb.CreateWishListReq{
		WishList: &pb.Wishlist{
			Id:   generateID(),
			Name: "my wishlist",
		},
	})

	if err != nil {
		panic("cannot connect with server " + err.Error())
	}

	fmt.Println("ID", res.WishListId)
}
