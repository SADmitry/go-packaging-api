FROM golang:1.24

WORKDIR /app

COPY . .

RUN go build -o go-packaging-api

EXPOSE 8080

CMD ["/app/go-packaging-api"]
