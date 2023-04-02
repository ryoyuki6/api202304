# https://hub.docker.com/_/golang
FROM golang:1.18.1-alpine
WORKDIR /app
# main.go と go.mod と go.sum を /app にコピー
COPY ["main.go", "go.mod", "go.sum", "./"]
# # パッケージをインストール
RUN go build
# # コンテナを起動する際に実行されるコマンド
CMD ["go", "run", "main.go"]