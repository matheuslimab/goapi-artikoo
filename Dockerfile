FROM golang:1.20 AS build

WORKDIR /app

COPY . /app

RUN go build -o server

FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /app/server /server

EXPOSE 8080

ENTRYPOINT [ "/server" ]