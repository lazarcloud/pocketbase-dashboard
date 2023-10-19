FROM golang:alpine as builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/main .

FROM alpine:latest
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]