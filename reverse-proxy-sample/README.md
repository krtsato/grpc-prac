# gRPC Reverse Proxy Sample

Go + RESTful API + gRPC に入門する

- curl コマンドで RESTful API にクエリを発行する
- RESTful API は gRPC Client となり gRPC Server へプロキシする
- gRPC Server はユーザ取得 / 作成 / 更新 / 削除 の結果を返却する
- 以下のユーザ情報を扱う

| RPC 名     | Method | URL                          | 役割           |
| ---------- | ------ | ---------------------------- | -------------- |
| ListUsers  | GET    | /api/v1/users                | ユーザ一覧取得 |
| GetUser    | GET    | /api/v1/users/{encrypted_id} | ユーザ一人取得 |
| CreateUser | POST   | /api/v1/users                | ユーザ作成     |
| UpdateUser | PUT    | /api/v1/users/{encrypted_id} | ユーザ更新     |
| DeleteUser | DELETE | /api/v1/users/{encrypted_id} | ユーザ削除     |

## 手順

- Dockerfile を書く
- `go mod init grpc-samples/reverse-proxy-sample` で Go Modules 化する
- サービス user に対して api/userpb/user.proto ファイルを書く
- 各種コードを自動生成する
  - protocol buffer となる api/userpb/user.pb.go を生成する
  - .proto ファイルの可読性を考慮して api/userpb/user.proto.html を生成する
    - markdown 形式は出力コードが汚いため忌避
  - リバースプロキシを提供する api/userpb/user.pb.gw.go を生成する
  - RESTful API を定義する api/userpb/user.swagger.json を生成する

```zsh
% pwd # ~/grpc-samples/reverse-proxy-sample

# 以下コンテナ内で実行
% docker-compose up --build -d
% protoc -I ./api -I /usr/local/include \
  -I $GOPATH/pkg/mod/github.com/grpc-ecosystem/grpc-gateway@v1.14.6/third_party/googleapis \
  --go_out=plugins=grpc:. \
  --doc_out=html,user.proto.html:./api/userpb \
  --grpc-gateway_out=logtostderr=true:. \
  --swagger_out=logtostderr=true:./api \
  ./api/userpb/user.proto
```

- cmd/server/main.go を書く
- cmd/client/main.go を書く
- cmd/proxy/main.go を書く
- 通信結果を確認

```zsh
% go run cmd/server/main.go
% go run cmd/client/main.go

% curl -X GET http://localhost:9997/api/v1/users # 一覧取得
% curl -X GET http://localhost:9997/api/v1/users/{encrypted_id} # 一人取得
% curl -X POST http://localhost:9997/api/v1/users -d {"name":"suzumiya"} # 作成
% curl -X PUT http://localhost:9997/api/v1/users/{encrypted_id} -d {"name":"haruhi"} # 更新
% curl -X DELETE http://localhost:9997/api/v1/users/{encrypted_id}  # 削除
```

- `go mod tidy` で依存パケージを適宜整理する
- `docker-compose down --rmi all` で終了する
