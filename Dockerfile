FROM golang:1.20.5-alpine3.17 AS builder

RUN go version

COPY . /goreminder
WORKDIR /goreminder

RUN go mod download
RUN go run github.com/steebchen/prisma-client-go generate
RUN go build -o ./.bin/app ./main.go

#lightweight docker container with binary
FROM node:20-alpine3.17

WORKDIR /root

COPY --from=builder /goreminder/.bin/app .
COPY --from=builder /goreminder/package.json .
COPY --from=builder /goreminder/package-lock.json .
COPY --from=builder /goreminder/migrations ./migrations
COPY --from=builder /goreminder/schema.prisma .

RUN node -v
RUN npm -v
RUN npm install


EXPOSE 3500

CMD npm run migrate && ./app