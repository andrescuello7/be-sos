FROM goland:last AS builder

RUN echo "Welcome"

RUN apt-get update \ echo "Finish"

WORKDIR /GoMessage

RUN go mod download \ go install \ go build

ENTRYPOINT ['./go-sms']
