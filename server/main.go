package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	pb "github.com/fxivan/go-grpc/proto"
)

type server struct {
	pb.UnimplementedWishListServiceServer //raper de nuestro server
}

func (s *server) Create(ctx context.Context, req *pb.CreateWishListReq) (*pb.CreateWishListRes, error) {
	fmt.Println("createing the wishlist " + req.WishList.Name)
	return &pb.CreateWishListRes{
		WishListId: req.WishList.Id,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":5051")
	if err != nil {
		panic("cannot create tcp connection " + err.Error())
	}

	serv := grpc.NewServer()
	pb.RegisterWishListServiceServer(serv, &server{})
	if err = serv.Serve(listener); err != nil {
		panic("cannot initialize the server" + err.Error())
	}

}
