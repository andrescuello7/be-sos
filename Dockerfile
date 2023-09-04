FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o my-go-app

RUN echo "Welcome"

RUN ls -l -a \ echo "Medium"

RUN echo "Finish"

CMD ["./my-go-app"]
