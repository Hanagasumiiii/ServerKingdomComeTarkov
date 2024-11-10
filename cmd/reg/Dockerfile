FROM golang:1.23.3-alpine AS builder

RUN apk update && apk add --no-cache git
RUN go install entgo.io/ent/cmd/ent@latest

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY ../../reg .

RUN go generate ./ent

RUN go build -o reg_service ./reg/main.go

FROM alpine:3.20.3

COPY --from=builder /app/reg_service ./

COPY --from=builder /app/ent /app/ent

EXPOSE 8080

CMD ["./reg_service"]

