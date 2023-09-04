FROM goland:last AS builder

RUN echo "Welcome"

WORKDIR /GoMessage

RUN ls -l -a \ echo "Medium"

RUN go mod download \ go install \ go build

RUN ls -l -a \ echo "Finish"

ENTRYPOINT ['./go-sms']
