FROM golang:latest

WORKDIR /go/src/app
COPY main .
COPY imgs ./imgs

CMD ["./main"]