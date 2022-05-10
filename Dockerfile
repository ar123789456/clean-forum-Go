FROM golang:latest
COPY ./ ./
RUN go build -o main ./cmd/api/main.go
CMD ["./main"]