FROM golang:1.18.3 AS builder

ENV GOOS=linux \
    CGO_ENABLED=0 \
    GO111MODULE=on \
    GOPROXY=https://goproxy.cn

COPY ./src /src
WORKDIR /src
RUN go build -o /build/app  .

FROM alpine:3.16.0

COPY ./env /env
COPY --from=builder /build/app /bin/app
CMD /bin/app