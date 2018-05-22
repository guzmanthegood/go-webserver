# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:latest
ENV SRC_DIR=/go/src/go-webserver/

# Add the source code:
COPY . $SRC_DIR
WORKDIR $SRC_DIR

# Build it:
RUN go get -v github.com/sirupsen/logrus
RUN export PATH=$PATH:$GOPATH/bin
RUN go build
ENTRYPOINT ["./go-webserver"]
EXPOSE 3000