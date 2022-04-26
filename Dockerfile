# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

# AWS_ACCESS_KEY_ID and AWS_SECRET_ACCESS_KEY must be set as --build-arg when
# running docker build so that they can be accessed by Dockerfile
ARG AWS_ACCESS_KEY_ID
ARG AWS_SECRET_ACCESS_KEY

ENV AWS_ACCESS_KEY_ID=$AWS_ACCESS_KEY_ID
ENV AWS_SECRET_ACCESS_KEY=$AWS_SECRET_ACCESS_KEY

WORKDIR /go-aws-pg

COPY . ./
RUN go mod download

RUN go build -o ./bin/server ./main.go

EXPOSE 8080

CMD [ "bin/server" ]
