FROM golang:1.15

MAINTAINER Maintainer

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src/gin_docker

COPY . /go/src/gin_docker

RUN go get -d -v ./...
RUN go install -v ./...

RUN go build gin_docker/src/app

EXPOSE $PORT

#ENTRYPOINT ["./app"]
