# syntax=docker/dockerfile:1

FROM golang:1.17 AS build
WORKDIR /lambda-go
COPY . ./
RUN go mod download
RUN CGO_ENABLED=0 go build -o server

FROM public.ecr.aws/c2t6n2x5/serverlessish:2 AS s

FROM gcr.io/distroless/static
COPY --from=s /opt/extensions/serverlessish /opt/extensions/serverlessish
COPY --from=build /lambda-go/server ./
COPY --from=build /lambda-go/.env ./

CMD ["/server"]
