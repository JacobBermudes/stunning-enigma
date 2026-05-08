FROM golang:1.25.5-alpine AS builder

WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY main.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o surfboost-server main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /root/

COPY --from=builder /app/surfboost-server .

COPY about.html .
COPY privacy.html .
COPY terms.html .
COPY logo.webp .

EXPOSE 8082

CMD ["./surfboost-server"]