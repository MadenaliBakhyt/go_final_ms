FROM golang:1.19-alpine

WORKDIR /app/teachers_service
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/teachers_service/main.go"]