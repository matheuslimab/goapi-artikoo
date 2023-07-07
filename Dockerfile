FROM golang:1.20 as builderApp
RUN mkdir app/
ADD . /app
WORKDIR /app

RUN go build -o main .

FROM nginx
COPY --from=builderApp /app/main /main
CMD ["/main"]