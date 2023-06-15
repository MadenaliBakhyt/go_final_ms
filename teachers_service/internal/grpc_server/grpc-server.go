package server

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"sync"
	"teachers_service/internal/data"
	"teachers_service/internal/db"
	"teachers_service/internal/validator"
	pb "teachers_service/pb"
	"teachers_service/pkg/config"
	"teachers_service/pkg/jsonlog"
)

type Server struct {
	pb.UnimplementedTeachersServiceServer
	Models data.Models
	Wg     sync.WaitGroup
	Logger *jsonlog.Logger
}

func (s *Server) GetTeacher(ctx context.Context, r *pb.TeacherId) (*pb.Teacher, error) {
	teacher, err := s.Models.Teachers.Get(r.GetId())
	if err != nil {
		return nil, status.Errorf(codes.NotFound, "Teacher not found")
	}
	return &pb.Teacher{Id: teacher.ID, Name: teacher.Name, Surname: teacher.Surname, Subject: teacher.Subject, Salary: teacher.Salary}, nil
}

func (s *Server) GetTeachers(context.Context, *pb.GetTeachersRequest) (*pb.TeacherList, error) {
	list, err := s.Models.Teachers.SelectAll()
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can not get list")
	}
	var teachers []*pb.Teacher

	for i := 0; i < len(list); i++ {
		teachers = append(teachers, &pb.Teacher{
			Id:      list[i].ID,
			Name:    list[i].Name,
			Surname: list[i].Surname,
			Subject: list[i].Subject,
			Salary:  list[i].Salary,
		})
	}
	return &pb.TeacherList{List: teachers}, nil
}

func (s *Server) CreateTeacher(ctx context.Context, r *pb.Teacher) (*pb.Teacher, error) {
	newTeacher := &data.Teacher{Name: r.Name, Surname: r.Surname, Subject: r.Subject, Salary: r.Salary}
	v := validator.New()
	if data.ValidateTeacher(v, newTeacher); !v.Valid() {
		return nil, status.Errorf(codes.InvalidArgument, "Invalid data")
	}
	err := s.Models.Teachers.Insert(newTeacher)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't insert data")
	}
	return r, nil
}

func (s *Server) DeleteTeacher(ctx context.Context, r *pb.TeacherId) (*pb.Teacher, error) {
	err := s.Models.Teachers.Delete(r.Id)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Can't delete teacher or not found")
	}
	return &pb.Teacher{Id: r.Id}, nil
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

	pb.RegisterTeachersServiceServer(server, &Server{
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
