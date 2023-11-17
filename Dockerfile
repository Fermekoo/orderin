# Build stage

FROM golang:1.19-alpine3.17 as builder
WORKDIR /app

COPY go.mod go.sum ./

COPY . .

#build go binary
RUN go build -o orderinapi cmd/main.go

RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.15.2/migrate.linux-amd64.tar.gz | tar xvz

#Run stage

FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/orderinapi .
COPY --from=builder /app/migrate ./migrate
COPY app.env ./app.env
COPY db/migrations ./migrations
COPY start.sh .



EXPOSE 8083

CMD [ "/app/orderinapi" ]

ENTRYPOINT [ "/app/start.sh" ]