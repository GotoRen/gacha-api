#Builder Container
FROM golang:1.15 AS builder

#環境変数をセット
ENV GOOS=linux
ENV GOARCH=amd64
ENV CGO_ENABLED=0
ENV GO111MODULE on

#作業ディレクトリを/appに指定
WORKDIR /app

#先にgo moduleをcopy
COPY go.mod .
COPY go.sum .

#go moduleのダウンロード（airは除去）
RUN go mod download

#api.game/を/appにcopy
COPY . .

#ビルド
RUN go build -o main

#マルチステージビルド
FROM alpine:3.11.6 as production

#オブジェクトファイルを/app/mainにcopy
COPY --from=builder /app/main .

#./mainを実行
CMD ["./main"]