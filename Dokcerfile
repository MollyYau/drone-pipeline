#
# ----- Go Builder Image ------
#
FROM golang:1.8-alpine AS builder

RUN apk add --no-cache git

# set working directory
RUN mkdir -p /go/src/hello
WORKDIR /go/src/hello

# copy sources
COPY . .

# add dependencies
RUN export http_proxy=http://192.168.2.155:2000 && \
  export https_proxy=http://192.168.2.155:2000 && \
  go get

# run tests
RUN go test -v

# build binary
RUN go build -v -o "/hello"

#
# ------ Hello World image ------
#

FROM alpine:3.6
MAINTAINER Ivan Pedrazas <ipedrazas@gmail.com>

LABEL description="Hello Go."
LABEL base="alpine"

COPY --from=builder /hello /bin/hello

ENTRYPOINT [ "/bin/hello" ]