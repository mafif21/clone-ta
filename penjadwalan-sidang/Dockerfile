FROM golang:1.21-alpine as builder

WORKDIR /app

COPY . .

ENV DB_HOST=mysql-container
ENV DB_USER=root
ENV DB_PASSWORD=root
ENV DB_NAME=penjadwalan_sidang
ENV DB_PORT=3306

ENV PORT=3000

ENV JWT_KEY=secret

RUN go mod download && go mod verify

RUN go mod tidy && go mod vendor

RUN go build -o binary cmd/main.go


FROM alpine:3.17

WORKDIR /app

COPY --from=builder /app .

EXPOSE 3000

CMD ["./binary"]