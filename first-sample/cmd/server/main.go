package main

import (
	"context"
	"errors"
	cat "grpc-samples/first-sample/api/catpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type myCatService struct{}

func (s *myCatService) GetMyCat(ctx context.Context, message *cat.GetMyCatMessage) (*cat.MyCatResponse, error) {
	switch message.TargetCat {
	case "tama":
		return &cat.MyCatResponse{
			Name: "tama",
			Kind: "Maine Coon",
		}, nil
	case "mike":
		return &cat.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	default:
		return nil, errors.New("Not found your cat.")
	}
}

func main() {
	port, err := net.Listen("tcp", ":9999")
	if err != nil {
		log.Fatal("listening port error: ", err.Error())
		return
	}

	s := grpc.NewServer()
	cat.RegisterCatServer(s, &myCatService{})
	s.Serve(port)
}
