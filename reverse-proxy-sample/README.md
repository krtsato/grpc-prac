# gRPC First Sample

## 手順

- Dockerfile を書く
- `go mod init grpc-samples/reverse-proxy-sample` Go Modules 化する
- サービス user に対して api/userpb/user.proto ファイルを書く
- protocol buffer となる api/userpb/user.pb.go を出力する

```zsh
% pwd # ~/grpc-samples/reverse-proxy-sample

# 開発環境 (ホストマシン上 or コンテナ内) で実行
% protoc --go_out=plugins=grpc:. api/userpb/user.proto
```

- リバースプロキシのコードを出力する

```zsh
% protoc --grpc-gateway_out=logtostderr=true:. api/userpb/user.proto
```

- api/*.proto を user.html として docs 配下に生成する

```zsh
% protoc --doc_out=html,user.html:docs api/userpb/*.proto
```

- cmd/server/main.go を書く
- cmd/client/main.go を書く
- 通信結果を確認

```zsh
% go run cmd/server/main.go
% go run cmd/client/main.go
```
