FROM golang:1.22

WORKDIR /opt/app

COPY go.* .

RUN go mod download

COPY . .

RUN go build cmd/users/main.go

EXPOSE 8020

CMD ["./main"]