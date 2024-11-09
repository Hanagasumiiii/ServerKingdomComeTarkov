FROM golang:1.23.3-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY . .

RUN go generate ./ent

RUN go build -o reg_service ./reg/main.go

FROM alpine:3.20.3

COPY --from=builder /app/reg_service ./

COPY --from=builder /app/ent /app/ent

EXPOSE 8080

CMD ["./reg_service"]

