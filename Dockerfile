# Dockerfile
FROM golang:1.21

WORKDIR /app
COPY server/ .

RUN go build -o server .

EXPOSE 8080
CMD ["./server"]
