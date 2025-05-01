FROM golang:1.24

WORKDIR /app
COPY server/ .
COPY data/ ../data/
COPY dist/ ./dist

RUN go build -o server .

EXPOSE 8080
CMD ["./server"]
