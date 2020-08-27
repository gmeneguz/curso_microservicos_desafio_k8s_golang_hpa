FROM golang:1.14.3-alpine3.11 AS builder

WORKDIR /go/src/app
COPY src ./src

RUN GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o /go/bin/main ./src/*.go

FROM scratch

COPY --from=builder /go/bin/main /go/bin/main
ENTRYPOINT ["/go/bin/main"]