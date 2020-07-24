# gRPC First Sample

## 手順

- Dockerfile を書く
- `go mod init grpc-samples/first-sample` Go Modules 化する
- サービス cat に対して api/catpb/cat.proto ファイルを書く
- protocol buffer となる api/catpb/cat.pb.go を出力する

```zsh
% pwd # ~/grpc-samples/first-sample

# 開発環境 (ホストマシン上 or コンテナ内) で実行
% protoc --go_out=plugins=grpc:. api/catpb/cat.proto
```

- api/*.proto を cat.html として docs 配下に生成する

```zsh
% protoc --doc_out=html,cat.html:docs api/catpb/*.proto
```

- cmd/server/main.go を書く
- cmd/client/main.go を書く
- 通信結果を確認

```zsh
% go run cmd/server/main.go
% go run cmd/client/main.go
```
