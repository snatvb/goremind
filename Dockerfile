FROM golang:1.20.5-alpine3.17 AS builder

RUN go version
RUN apk add git

COPY ./ /snatvb/goreminder
WORKDIR /snatvb/goreminder

RUN go mod download && go get -u ./...
RUN CGO_ENABLED=0 GOOS=linux go build -o ./.bin/app ./cmd/api/main.go

#lightweight docker container with binary
FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=0 /snatvb/goreminder/.bin/app .

EXPOSE 3500

CMD [ "./app"]