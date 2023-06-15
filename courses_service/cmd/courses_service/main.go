package main

import (
	server "courses_service/internal/grpc_server"
	"courses_service/internal/rabbit"
)

func main() {
	server.Start()
	rabbit.StartRabbitServer()
}
