FROM golang:1.11.3-alpine3.8 as builder

COPY . /go/src/github.com/mslusarczyk/median
WORKDIR /go/src/github.com/mslusarczyk/median

RUN go test ./...
RUN go build -o median .

FROM alpine:3.8

RUN adduser -S -D -H -h /app appuser
USER appuser

COPY --from=builder /go/src/github.com/mslusarczyk/median/median /app/
WORKDIR /app

ENTRYPOINT ["./median"]