FROM golang:alpine

RUN mkdir app

WORKDIR /app

Add . /app

RUN go build -o server .

CMD ["./server"]