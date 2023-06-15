package main

import (
	server "teachers_service/internal/grpc_server"
	"teachers_service/internal/rabbit"
)

func main() {
	server.Start()
	rabbit.StartRabbitServer()
}
