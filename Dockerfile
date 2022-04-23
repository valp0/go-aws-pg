# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /go-aws-pg

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY * ./

RUN go build -o /go-aws-pg

EXPOSE 8080

CMD [ "/go-aws-pg" ]
