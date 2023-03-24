FROM golang:1.20 as builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o artikooApi 

FROM scratch
COPY --from=builder /app/artikooApi /artikooApi

ENTRYPOINT ["/artikooApi"]