FROM ubuntu:18.04
USER root

# golang general
RUN apt-get update && apt-get install -y golang-go
VOLUME /opt/gopath
WORKDIR /opt/gopath
ENV GOPATH /opt/gopath

CMD go build github.com/Schneereich/merge-intervals