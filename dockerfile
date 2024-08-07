FROM golang:1.22-alpine

WORKDIR /app

COPY . .

RUN go mod download


RUN go build -o my-go-app

EXPOSE 8081

CMD ["./my-go-app"]
