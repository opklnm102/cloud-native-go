FROM golang:1.9.2-alpine3.7
MAINTAINER Ethen Kim(opklnm102@gmail.com)

ENV SOURCES /go/src/cloud-native-go

COPY . ${SOURCES}
RUN cd ${SOURCES} && CGO_ENABLED=0 go install
ENV PORT 8080
EXPOSE 8080

ENTRYPOINT cloud-native-go
