package main

import (
	"context"
	cat "first-sample/api"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9999", grpc.WithInsecure())
	if err != nil {
		log.Fatal("connection error: ", err)
	}
	defer conn.Close()

	client := cat.NewCatClient(conn)
	message := &cat.GetMyCatMessage{TargetCat: "mike"}
	res, err := client.GetMyCat(context.Background(), message)
	if err != nil {
		log.Fatal("getting cat error: ", err)
	}

	fmt.Printf("result:%s\n", res)
}
