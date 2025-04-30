FROM golang:1.24

WORKDIR /app
COPY server/ .
COPY data/ ../data/

RUN go build -o server .

EXPOSE 8080
CMD ["./server"]
