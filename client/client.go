package main

import (
	"context"
	"log"
	"time"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

func main() {

	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// GetStudent
	res, err := client.GetStudent(ctx, &pb.StudentRequest{Id: 101})
	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Println("Student Info:")
	log.Println("ID:", res.Id)
	log.Println("Name:", res.Name)
	log.Println("Major:", res.Major)
	log.Println("Email:", res.Email)
	log.Println("Phone:", res.Phone)

	// ListStudents
	log.Println("\nStudent List:")

	listRes, err := client.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}

	for _, s := range listRes.Student {

		log.Println("ID:", s.Id)
		log.Println("Name:", s.Name)
		log.Println("Major:", s.Major)
		log.Println("Email:", s.Email)
		log.Println("Phone:", s.Phone)
		log.Println("------")
	}

}