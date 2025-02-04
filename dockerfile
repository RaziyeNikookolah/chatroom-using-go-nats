FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN GO111MODULE=on GOPROXY=https://goproxy.cn,direct go mod download
COPY . .
RUN GO111MODULE=on GOPROXY=https://goproxy.cn,direct go build -o server ./cmd/chatroom/main.go

FROM alpine:latest
RUN apk --no-cache add tzdata
ENV TZ=Asia/Tehran
RUN apk --no-cache add ca-certificates
WORKDIR /app/
COPY --from=builder /app/server .
EXPOSE 50051
CMD ["./server"]