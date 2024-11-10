FROM golang:1.23.3-alpine AS builder

RUN apk update && apk add --no-cache git
RUN go install entgo.io/ent/cmd/ent@latest

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go generate ./ent

RUN go build -o auth_service ./cmd/auth/main.go

FROM alpine:3.20.3

WORKDIR /root/

COPY --from=builder /app/auth_service ./

COPY --from=builder /app/ent /app/ent

EXPOSE 8081

CMD ["./auth_service"]