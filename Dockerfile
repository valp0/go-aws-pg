# syntax=docker/dockerfile:1

FROM golang:1.17-alpine
WORKDIR /go-aws-pg

COPY . ./
RUN go mod download

RUN go build -o ./bin/server ./main.go

CMD [ "bin/server" ]
