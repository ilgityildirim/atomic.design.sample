FROM golang:1.21-alpine3.19
MAINTAINER "Ilgıt Yıldırım <ilgit.yildirim@triplebits.com>"

WORKDIR /app

COPY ./cmd/go.mod ./cmd/go.sum ./
RUN go mod download

#USER 1000:1000
