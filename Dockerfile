FROM goland:last AS builder

RUN apt-get update \ ./deploy

WORKDIR /GoMessage

RUN go mod download \ go install \ go build

ENTRYPOINT ['./go-sms']
