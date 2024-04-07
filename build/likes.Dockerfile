FROM golang:1.22

WORKDIR /opt/app

COPY go.* .

RUN go mod download

COPY . .

RUN go build cmd/likes/main.go

EXPOSE 8030

CMD ["./main"]