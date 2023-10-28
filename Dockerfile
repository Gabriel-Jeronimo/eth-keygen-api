FROM golang:1.21.3-bullseye

COPY .env /app/.env

WORKDIR /app

COPY *.go ./

COPY . . 
EXPOSE 8080

CMD ["go", "run", "./cmd/api/main.go"]