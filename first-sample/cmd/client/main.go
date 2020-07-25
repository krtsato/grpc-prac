package main

import (
	"context"
	"fmt"
	catpb "grpc-samples/first-sample/api/catpb"
	"log"

	"google.golang.org/grpc"
)

const address = "localhost:9999"

func run() error {
	// gRPC コネクション作成
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		return err
	}
	defer conn.Close()

	// client 側の通信準備
	client := catpb.NewCatClient(conn)
	message := &catpb.GetMyCatMessage{TargetCat: "mike"}

	// データ取得
	res, err := client.GetMyCat(context.Background(), message)
	if err != nil {
		return err
	}

	fmt.Printf("Result: %s\n", res)
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatalln(err.Error())
	}
}
