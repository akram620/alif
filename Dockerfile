FROM golang:1.22

WORKDIR /app

# Загружаем зависимости
COPY ["go.mod", "go.sum", "./"]
RUN go mod download

COPY ./ ./

RUN go build -o main cmd/server/main.go

EXPOSE 8084

CMD ["./main"]
