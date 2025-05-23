# services/accounts/Dockerfile

FROM golang:1.21-alpine

WORKDIR /app
COPY . .

RUN go mod tidy
RUN go build -o accounts-service .

EXPOSE 8080
CMD ["./accounts-service"]
