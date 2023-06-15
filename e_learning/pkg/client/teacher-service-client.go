package client

import (
	"e-learning/internal/jsonlog"
	pb "e-learning/pb/teachers-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetTeachersClient() pb.TeachersServiceClient {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelError)
	dial, err := grpc.Dial("localhost:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//dial, err := grpc.Dial("teachers:8003", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.PrintError(err, nil)
	}
	return pb.NewTeachersServiceClient(dial)
}
