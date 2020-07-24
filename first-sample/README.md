# gRPC First Sample

## 手順

- Dockerfile を書く
- サービス cat に対して api/cat.proto ファイルを書く
- protocol buffer となる api/cat.pb.go を出力する

```zsh
% pwd # ~/grpc-samples/first-sample

# 開発環境 (ホストマシン上 or コンテナ内) で実行
% protoc --go_out=plugins=grpc:. api/cat.proto
```

- api/*.proto を cat.html として docs 配下に生成する

```zsh
% protoc --doc_out=html,cat.html:docs api/*.proto
```

- cmd/server/main.go を書く
- cmd/client/main.go を書く
