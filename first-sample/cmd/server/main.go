package main

import (
	"context"
	"errors"
	catpb "grpc-samples/first-sample/api/catpb"
	"log"
	"net"

	"google.golang.org/grpc"
)

const port = ":9999"

type myCatService struct{}

// Context は API 境界を跨いだアクセスに用いる
// e.g. キャンセルフラグ，リクエストごとの変数，リクエストのデッドライン
func (s *myCatService) GetMyCat(ctx context.Context, message *catpb.GetMyCatMessage) (*catpb.MyCatResponse, error) {
	switch message.TargetCat {
	case "tama":
		return &catpb.MyCatResponse{
			Name: "tama",
			Kind: "Maine Coon",
		}, nil
	case "mike":
		return &catpb.MyCatResponse{
			Name: "mike",
			Kind: "Norwegian Forest Cat",
		}, nil
	default:
		return nil, errors.New("Not found your cat.")
	}
}

func run() error {
	// Listen 状態を開始
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return err
	}

	// gRPC サーバを初期化・登録
	grpcServer := grpc.NewServer()
	catpb.RegisterCatServer(grpcServer, &myCatService{})

	// gRPC サーバを起動
	err = grpcServer.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err.Error())
	}
}
