FROM golang:alpine AS builder

LABEL stage=gobuilder

ENV CGO_ENABLED 0
ENV GOOS linux
ENV GOPROXY https://goproxy.cn,direct

WORKDIR /build/app

ADD go.mod .
ADD go.sum .
RUN go mod tidy
COPY . .
RUN go build -ldflags="-s -w" -o /app/app src/cmd/main.go

FROM alpine

RUN apk update --no-cache && apk add --no-cache ca-certificates tzdata
ENV TZ Asia/Shanghai

WORKDIR /app
COPY --from=builder /app/app /app/app

CMD ["./app"]
