FROM golang:latest

RUN mkdir -p /go/src/golang-docker
ADD . /go/src/golang-docker

WORKDIR /go/src/golang-docker

RUN go build -o /go/bin/golang-docker .
CMD ["/go/bin/golang-docker"]