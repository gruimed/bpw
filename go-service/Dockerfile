FROM golang:alpine3.20

WORKDIR /go/

COPY ./src/* /go/

RUN go mod download
RUN go build -o dice
ENTRYPOINT ["/go/dice"]

