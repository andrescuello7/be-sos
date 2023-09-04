FROM goland:last AS builder
RUN apt-get update
ENV URL_PATH=https://api.twilio.com/2010-04-01/Accounts/ACe6878cfeb48ec451d63d86224bf7e19c/Messages.json \
    ACCOUNT_SID=ACe6878cfeb48ec451d63d86224bf7e19c \
    AUTH_TOKEN=196d8a5d340f225c51f4e2b6db797177 \
    TO_PHONE=5493815615474 \
    FROM_PHONE=14155238886 \
WORKDIR /go/src
COPY go.mod .
RUN go mod download
COPY . .
RUN go install
RUN go build
ENTRYPOINT ['./go-sms']
