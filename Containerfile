FROM golang:1.23-alpine AS builder
RUN export GOBIN=$HOME/work/bin
WORKDIR /go/src/app
ADD . .
RUN go get -v ./...
RUN CGO_ENABLED=0 go build -ldflags="-w -s" -o main .

FROM alpine:3
RUN apk add --no-cache openssh-client sshpass ca-certificates curl
COPY --from=builder /go/src/app/main /app/
ADD demo.sh /app/
WORKDIR /app
CMD ["./main"]