FROM golang:1.10-alpine
WORKDIR /go/src/github.com/julianespinel/mila-api
COPY . .
RUN go build -o goapp
ENTRYPOINT ./goapp
