# gRPC First Sample

Go + gRPC に入門する  

- 各種コマンドと出力結果を確認する
- Client が猫の名前をサーバに送る
- Server は猫の名前と種類を返却する

## 手順

- Dockerfile を書く
- `go mod init grpc-samples/first-sample` Go Modules 化する
- サービス cat に対して api/catpb/cat.proto ファイルを書く
- protocol buffer となる api/catpb/cat.pb.go を出力する

```zsh
% pwd # ~/grpc-samples/first-sample

# 以下コンテナ内で実行
% docker-compose up --build -d
% protoc -I ./api --go_out=plugins=grpc:. api/catpb/cat.proto
```

- api/*.proto を cat.html として docs 配下に生成する
- markdown 形式は出力コードが汚いため忌避

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

- `go mod tidy` で依存パケージを適宜整理する
- `docker-compose down --rmi all` で終了する
