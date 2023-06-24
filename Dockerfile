FROM golang:1.20.5-alpine3.17 AS builder

RUN go version

COPY . /goreminder
WORKDIR /goreminder

RUN go mod download
RUN go run github.com/steebchen/prisma-client-go db push
RUN go build -o ./.bin/app ./main.go

#lightweight docker container with binary
FROM alpine:latest

WORKDIR /root/

COPY --from=0 /goreminder/.bin/app .

EXPOSE 3500

CMD [ "./app"]