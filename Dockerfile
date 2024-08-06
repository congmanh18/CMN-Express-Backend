FROM golang:1.22-alpine AS builder

WORKDIR /project

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o app cmd/main.go

FROM alpine:3.18

WORKDIR /dist

COPY --from=builder /project/app .
COPY wait-for-it.sh .

RUN chmod +x wait-for-it.sh \
    && apk add --no-cache postgresql-client

CMD ["./wait-for-it.sh", "db", "./app"]
