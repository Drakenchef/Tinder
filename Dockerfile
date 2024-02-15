FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY  . .
RUN go build -o TheTinder cmd/main/main.go

FROM alpine as run_stage
WORKDIR /out
COPY --from=builder /app/TheTinder ./binary
EXPOSE 8000
CMD ["./binary"]