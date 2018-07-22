FROM golang:1.10-alpine
RUN apk add --update bash && rm -rf /var/cache/apk/*
WORKDIR /go/src/github.com/julianespinel/mila-api
COPY . .
RUN go build -o goapp
