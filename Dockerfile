FROM golang:1.20 as builderApp
RUN mkdir app/
ADD . /app
WORKDIR /app

RUN GOOS=linux GOARCH=amd64  go build -o main .

FROM scratch
COPY --from=builderApp /app/main /main
# CMD ["/main"]
ENTRYPOINT [ "/main" ]