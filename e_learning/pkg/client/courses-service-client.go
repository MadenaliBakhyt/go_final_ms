package client

import (
	"e-learning/internal/jsonlog"
	pb "e-learning/pb/courses-service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"os"
)

func GetCoursesClient() pb.CoursesServiceClient {
	logger := jsonlog.New(os.Stdout, jsonlog.LevelError)
	dial, err := grpc.Dial("localhost:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	//dial, err := grpc.Dial("courses:8002", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logger.PrintError(err, nil)
	}
	return pb.NewCoursesServiceClient(dial)
}
