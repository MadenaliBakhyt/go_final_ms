package server

import (
	"context"
	"courses_service/internal/data"
	"courses_service/internal/db"
	"courses_service/internal/validator"
	pb "courses_service/pb"
	"courses_service/pkg/config"
	"courses_service/pkg/jsonlog"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"sync"
)

type Server struct {
	pb.UnimplementedCoursesServiceServer
	Models data.Models
	Wg     sync.WaitGroup
	Logger *jsonlog.Logger
}

func (s *Server) GetCourse(ctx context.Context, r *pb.CourseId) (*pb.Course, error) {
	course, err := s.Models.Courses.Get(r.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Course not found")
	}
	return &pb.Course{Id: course.ID, Name: course.Name, Credits: course.Credits, Price: course.Price}, nil
}

func (s *Server) GetCourses(context.Context, *pb.GetCoursesRequest) (*pb.CourseList, error) {
	list, err := s.Models.Courses.SelectAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can not get list")
	}
	var coursesAll []*pb.Course

	for i := 0; i < len(list); i++ {
		coursesAll = append(coursesAll, &pb.Course{
			Id:      list[i].ID,
			Name:    list[i].Name,
			Credits: list[i].Credits,
			Price:   list[i].Price,
		})
	}
	return &pb.CourseList{List: coursesAll}, nil
}

func (s *Server) CreateCourse(ctx context.Context, r *pb.Course) (*pb.Course, error) {
	newCourse := &data.Course{Name: r.Name, Credits: r.Credits, Price: r.Price}
	v := validator.New()
	if data.ValidateCourse(v, newCourse); !v.Valid() {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid data")
	}
	err := s.Models.Courses.Insert(newCourse)
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return r, nil
}

func (s *Server) DeleteCourse(ctx context.Context, r *pb.CourseId) (*pb.Course, error) {
	err := s.Models.Courses.Delete(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't delete course, not found")
	}
	return &pb.Course{Id: r.Id}, nil
}

func Start() {

	conf := config.GetConfig()

	logger := jsonlog.New(os.Stdout, jsonlog.LevelInfo)

	db, err := db.OpenDB()

	if err != nil {
		logger.PrintFatal(err, nil)
	}

	logger.PrintInfo("database connection pool established", nil)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", conf.Port))

	if err != nil {
		log.Fatalf("failed to listen: %s", err)
	}

	server := grpc.NewServer()

	reflection.Register(server)

	pb.RegisterCoursesServiceServer(server, &Server{
		Logger: logger,
		Models: data.NewModels(db),
	})

	log.Printf("Server listening at %d", conf.Port)

	//server.Serve(lis)
	go serveGoroutine(server, lis)
}

func serveGoroutine(s *grpc.Server, lis net.Listener) {
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
