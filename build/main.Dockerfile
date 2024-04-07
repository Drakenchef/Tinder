FROM golang:1.22

WORKDIR /opt/app

COPY go.* .

RUN go mod download

COPY . .

RUN go build cmd/main/main.go

EXPOSE 8000

CMD ["./main"]