# Builder
FROM golang:1.13.4-alpine3.10 as builder

RUN apk update && apk upgrade && \
  apk --update add git make

WORKDIR /app

COPY . .

RUN go build -o events-api main.go

# Distribution
FROM alpine:latest

RUN apk update && apk upgrade && \
  apk --update --no-cache add tzdata && \
  mkdir /app

EXPOSE 9090

COPY --from=builder app/events-api /app

CMD app/events-api