# 1. ビルドステージ
FROM golang:1.24.3 AS builder

WORKDIR /app

COPY . .

RUN go mod download

# SQLite3用にCGOを有効化
RUN apt-get update && apt-get install -y gcc

RUN CGO_ENABLED=1 GOOS=linux go build -o server ./cmd/server

RUN ls -l /app/ # ビルドされたバイナリを確認
RUN ls -lh /app/server && file /app/server # バイナリの詳細を確認
CMD ["/bin/sh", "-c", "echo 🔧 launching... && /app/server"]

# 2. 実行ステージ（Golangに合わせた環境）
FROM debian:bookworm-slim

WORKDIR /app

COPY --from=builder /app/server /app/server
COPY --from=builder /app/papers.db /tmp/papers.db
RUN ls -l /tmp/

COPY --from=builder /etc/ssl/certs /etc/ssl/certs

EXPOSE 8080

CMD ["/app/server"]
