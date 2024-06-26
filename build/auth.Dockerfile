FROM golang:1.22

WORKDIR /opt/app

COPY go.* .

RUN go mod download

COPY . .

RUN go build cmd/auth/main.go

EXPOSE 8010

CMD ["./main"]