FROM golang:1.23 as builder

WORKDIR /github.com/RaziyeNikookolah/chatroom-using-go-nats

RUN CGO_ENABLED=0 go build -o ./order-api cmd/main.go

FROM alpine as deploy

RUN apk add --no-cache tzdata
ENV TZ=Asia/Tehran

WORKDIR /github.com/RaziyeNikookolah/chatroom-using-go-nats

COPY --from=builder /github.com/RaziyeNikookolah/chatroom-using-go-nats/order-api ./order-api

CMD [ "./order-api" , "--config" , "/etc/config.json"]