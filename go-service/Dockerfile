FROM golang:alpine3.20


#RUN apk add curl git make llvm clang 

#WORKDIR /go/src/

#RUN git clone https://github.com/open-telemetry/opentelemetry-go-instrumentation.git

#WORKDIR /go/src/opentelemetry-go-instrumentation/

#RUN make build

WORKDIR /go/

COPY ./src/* /go/

#RUN go mod init dice
RUN go get github.com/go-sql-driver/mysql
RUN go mod download
RUN go build -o dice
ENTRYPOINT ["/go/dice"]

#CMD /go/src/run.sh