FROM golang:1.23-alpine as builder
WORKDIR /app
COPY . .
RUN go mod tidy && go build -o main main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
EXPOSE 8080
CMD ["./main"]