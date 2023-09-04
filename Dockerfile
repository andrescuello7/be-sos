FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o my-go-app

CMD ["./my-go-app"]
