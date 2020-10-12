FROM golang:1.15

ENV GIN_MODE=release
ENV PORT=8080

WORKDIR /go/src/gin_docker

COPY . /go/src/gin_docker

RUN go get -d -v ./...
RUN go install -v ./...

# RUN go build gin_docker

RUN go get -u github.com/cosmtrek/air

EXPOSE $PORT