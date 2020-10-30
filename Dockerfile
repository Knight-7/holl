FROM golang:1.15.3:alpine as builder

WORKDIR /holl

COPY . /holl

RUN apk --no-cache add git

ENV GO111MODULE=on \
    CGO_ENABLE=0 \
    GOOS=linux

RUN go build -o app .

FROM alpine:latest

RUN apk --no-cache add ca-certificate

COPY --from=builder /holl/app .

CMD [ "./app" ]