FROM golang:1.15.3-alpine as builder

WORKDIR /holl

COPY . /holl

RUN apk --no-cache add git

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux \
    GOPROXY=https://mirrors.aliyun.com/goproxy/

RUN go build -o app .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /config

COPY config/application.yaml /config
COPY config/data.zip /config

WORKDIR /

COPY --from=builder /holl/app .

RUN chmod +x wait-for

CMD ["./app"]