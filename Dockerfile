FROM golang

ADD . /go/src/github.com/4others/nearest-location-server

RUN go get -d -v ./...
RUN go install github.com/4others/nearest-location-server

ENTRYPOINT ["/go/bin/nearest-location-server"]

EXPOSE 8080