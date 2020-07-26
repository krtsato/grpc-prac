package main

import (
	"flag"
	userpb "grpc-samples/reverse-proxy-sample/api/userpb"
	"log"
	"math/rand"
	"net"
	"time"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// DB から取得したものと仮定する
const (
	USERS = map[string]string{"12345abcde": "taro", "zxcvb09876": "hanako"}
	PORT  = "9998"
)

type userService struct{}

// ユーザ一覧取得
func (e *userService) ListUsers(ctx context.Context, req *userpb.ListUserRequest) (*userpb.ListUsersResponses, error) {
	var users = []*userpb.User{}
	for k, v := range USERS {
		users = append(users, &userpb.User{EncryptedId: k, Name: v})
	}

	return &userpb.ListUsersResponses{Users: users}, nil
}

// 1ユーザの取得(本来は詳細情報を付与して返す)
func (e *userService) GetUser(ctx context.Context, req *userpb.GetUserRequest) (*userpb.User, error) {
	return &userpb.User{EncryptedId: req.EncryptedId, Name: USERS[req.EncryptedId]}, nil
}

// ユーザの作成
func (e *userService) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.User, error) {
	encrypedId := GetRandString(10)
	USERS[encrypedId] = req.Name

	return &userpb.User{EncryptedId: encrypedId, Name: req.Name}, nil
}

// ユーザの更新
func (e *userService) UpdateUser(ctx context.Context, req *userpb.UpdateUserRequest) (*userpb.User, error) {
	USERS[req.EncryptedId] = req.Name

	return &userpb.User{EncryptedId: req.EncryptedId, Name: req.Name}, nil
}

// ユーザ削除
func (e *userService) DeleteUser(ctx context.Context, req *userpb.DeleteUserRequest) (*userpb.Empty, error) {
	delete(USERS, req.EncryptedId)
	return &userpb.Empty{}, nil
}

// 乱数の生成(encrypted_id用)
func GetRandString(n int) string {
	var letterRunes = []rune("1234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	var port string
	flag.StringVar(&port, "port", PORT, "")
	flag.Parse()

	listen, err := net.Listen("tcp", ":"+port)
	if err != nil {
		log.Fatalln(err)
	}

	server := grpc.NewServer()
	userpb.RegisterUserServiceServer(server, &userService{})
	reflection.Register(server)

	if err := server.Serve(listen); err != nil {
		log.Fatalln(err)
	}
}
