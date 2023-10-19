FROM golang:alpine as builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

# FROM alpine  do do, start with alpine and install docker compose
FROM docker:latest
# FROM alpine

# RUN apk update
# RUN apk add --no-cache docker-cli python3
# RUN apk add --no-cache --virtual .docker-compose-deps python3-dev libffi-dev openssl-dev gcc libc-dev make
# RUN pip3 install docker-compose
# RUN apk del .docker-compose-deps

# FROM ubuntu:20.04
USER root
# RUN apt update
# RUN apt install docker -y


COPY --from=builder /app/main .
# COPY ./data/exampleProject /app/data/projects/exampleProject
COPY ./clean-data ./clean-data
# COPY ./compose.sh /app/compose.sh
# WORKDIR /app
CMD ["./main"]