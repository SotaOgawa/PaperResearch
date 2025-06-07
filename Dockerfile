# 1. ビルドステージ
FROM golang:1.24.3 AS builder

WORKDIR /app

COPY . .

RUN go mod download

# SQLite3用にCGOを有効化
RUN apt-get update && apt-get install -y gcc

RUN CGO_ENABLED=1 GOOS=linux go build -o server ./cmd/server

RUN ls -l /app/ # ビルドされたバイナリを確認

# 2. 実行ステージ（小さくて安全）
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/server /app/server
COPY --from=builder /app/papers.db /app/papers.db # 必要に応じて
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

EXPOSE 8080

CMD ["/app/server"]
