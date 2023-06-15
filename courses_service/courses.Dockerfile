FROM golang:1.19-alpine

WORKDIR /app/courses_service
COPY . .

RUN go mod download

CMD ["go", "run", "cmd/courses_service/main.go"]