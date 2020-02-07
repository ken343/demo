FROM golang:rc-buster AS builder
RUN mkdir -p /home/go/src/github.com/ken343/demo/
ADD . /home/go/src/github.com/ken343/demo/
WORKDIR /home/go/src/github.com/ken343/demo/

RUN go build .

FROM debian:latest

COPY --from=builder /home/go/src/github.com/ken343/demo/ /home/servers/
WORKDIR /home/servers/
EXPOSE 80 443 22
EXPOSE 8081 8082 8083
EXPOSE 9090
RUN mv /home/servers/demo /bin/
CMD [ "demo" ]
