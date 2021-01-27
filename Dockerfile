FROM golang:alpine as builder

ENV GO111MODILE=on

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY go.mod ./
RUN go mod download
COPY . .
RUN go build -o main .

FROM alpine:latest
WORKDIR /usr/src/app
COPY --from=builder /app/main .

CMD ["./main"]