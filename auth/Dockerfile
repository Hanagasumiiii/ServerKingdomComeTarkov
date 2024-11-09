FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go generate ./ent

RUN go build -o auth_service ./auth/main.go

FROM alpine:3.20.3

WORKDIR /root/

COPY --from=builder /app/auth_service ./

COPY --from=builder /app/ent /app/ent

EXPOSE 8081

CMD ["./auth_service"]