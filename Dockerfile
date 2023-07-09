FROM golang:1.20 AS build

WORKDIR /app

COPY .env /app
COPY . /app

RUN go build -o server

FROM ubuntu:22.04

WORKDIR /

COPY --from=build /app/server /server

EXPOSE 8080

ENTRYPOINT [ "/server" ]