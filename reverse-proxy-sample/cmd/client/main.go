package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"

	userpb "grpc-samples/reverse-proxy-sample/api/userpb"
)

const HOST = "localhost:9998"

// 全ユーザ取得
func ListUser(conn *grpc.ClientConn) {
	client := userpb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.ListUsers(ctx, &userpb.ListUserRequest{})
	if err != nil {
		log.Fatalln(err)
	}

	for _, v := range resp.Users {
		log.Printf("EncryptedId: %s, Name: %s\n", v.EncryptedId, v.Name)
	}
}

// 1ユーザ取得
func GetUser(conn *grpc.ClientConn, encryptedId string) {
	client := userpb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.GetUser(ctx, &userpb.GetUserRequest{EncryptedId: encryptedId})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("EncryptedId: %s, Name: %s\n", resp.EncryptedId, resp.Name)
}

// ユーザ作成
func CreatUser(conn *grpc.ClientConn, name string) {
	client := userpb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.CreateUser(ctx, &userpb.CreateUserRequest{Name: name})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("EncryptedId: %s, Name: %s\n", resp.EncryptedId, resp.Name)
}

// ユーザ更新
func UpdateUser(conn *grpc.ClientConn, encryptedId string, name string) {
	client := userpb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	resp, err := client.UpdateUser(ctx, &userpb.UpdateUserRequest{EncryptedId: encryptedId, Name: name})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("EncryptedId: %s, Name: %s\n", resp.EncryptedId, resp.Name)
}

// ユーザ削除
func DeleteUser(conn *grpc.ClientConn, encryptedId string) {
	client := userpb.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	_, err := client.DeleteUser(ctx, &userpb.DeleteUserRequest{EncryptedId: encryptedId})
	if err != nil {
		log.Fatalln(err)
	}

	log.Printf("deleted")
}

// 標準入力受付
func GetStdInput(message string) string {
	fmt.Println(message)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	return scanner.Text()
}

func main() {
	var addr string
	flag.StringVar(&addr, "addr", HOST, "")
	flag.Parse()

	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

L:

	for {
		cmd := GetStdInput("Input ListUsers/GetUser/CreateUser/UpdateUser/DeleteUser/exit\n")

		switch cmd {
		case "ListUsers":
			ListUser(conn)
		case "GetUser":
			encryptedId := GetStdInput("encryptedId:")
			GetUser(conn, encryptedId)
		case "CreateUser":
			name := GetStdInput("name:")
			CreatUser(conn, name)
		case "UpdateUser":
			encryptedId := GetStdInput("encryptedId:")
			name := GetStdInput("name:")
			UpdateUser(conn, encryptedId, name)
		case "DeleteUser":
			encryptedId := GetStdInput("encryptedId:")
			DeleteUser(conn, encryptedId)
		case "exit":
			break L
		}
	}
}
