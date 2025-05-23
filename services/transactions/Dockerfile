FROM golang:1.21-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . ./

RUN go build -o transactions-service .

EXPOSE 8081

CMD ["./transactions-service"]
