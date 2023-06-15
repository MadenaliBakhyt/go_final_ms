FROM golang:1.19-alpine

WORKDIR /app/e_learning
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/api/main.go"]