FROM golang:1.10.3

ENV GOPATH /go
ENV GIN_MODE=release

ADD . /go/src/github.com/dimartiro/urlshort
WORKDIR /go/src/github.com/dimartiro/urlshort

# Build the binary
RUN go build -i main.go

# Expose port 8080
EXPOSE 8080

ENTRYPOINT ./main