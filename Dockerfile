FROM golang:1.10-alpine

WORKDIR /app
ADD . /app
RUN cd /app && go build -o goapp

ENTRYPOINT ./goapp
