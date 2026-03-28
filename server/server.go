package main

import (
	"context"
	"log"
	"net"

	pb "grpc-student/studentpb"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (s *server) GetStudent(ctx context.Context, req *pb.StudentRequest) (*pb.StudentResponse, error) {

	log.Printf("Received request for student ID: %d", req.Id)

	return &pb.StudentResponse{
		Id:    req.Id,
		Name:  "Alice Johnson",
		Major: "Computer Science",
		Email: "alice@university.com",
		Phone: "0812345678",
	}, nil
}

func (s *server) ListStudents(ctx context.Context, req *pb.Empty) (*pb.StudentListResponse, error) {

	log.Println("Received request for student list")

	students := []*pb.StudentResponse{
		{
			Id:    1,
			Name:  "Alice Johnson",
			Major: "Computer Science",
			Email: "alice@university.com",
			Phone: "0812345678",
		},
		{
			Id:    2,
			Name:  "Bob Smith",
			Major: "Information Technology",
			Email: "bob@university.com",
			Phone: "0898765432",
		},
	}

	return &pb.StudentListResponse{
		Student: students,
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	pb.RegisterStudentServiceServer(grpcServer, &server{})

	log.Println("gRPC Server running on port 50051")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}