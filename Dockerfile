ARG GOARCH=amd64
ARG GOOS=linux

FROM golang:1.16 AS builder

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct

WORKDIR /go/src/github.com/zhang21/learn_gin

COPY . .

RUN CGO_ENABLED=0 go build


FROM alpine:3.14

ENV GIN_MODE=release

WORKDIR /app

COPY --from=builder /go/src/github.com/zhang21/learn_gin/learn_gin /app/learn_gin
COPY ./conf /app/conf

EXPOSE 8000

CMD ["/app/learn_gin"]
