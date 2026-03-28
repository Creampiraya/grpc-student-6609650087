package main

import (
	"context"
	"encoding/json"
	"log"
	"time"

	pb "grpc-student/studentpb"

	"google.golang.org/grpc"
)

func main() {

	// Connect to gRPC server
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStudentServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	// ---------------- GetStudent ----------------
	res, err := client.GetStudent(ctx, &pb.StudentRequest{Id: 1})
	if err != nil {
		log.Fatalf("Error calling GetStudent: %v", err)
	}

	log.Println("Student Info:")
	log.Println("ID:", res.Id)
	log.Println("Name:", res.Name)
	log.Println("Major:", res.Major)
	log.Println("Email:", res.Email)
	log.Println("Phone:", res.Phone)

	// ---------------- ListStudents ----------------
	log.Println("\nStudent List (JSON):")

	listRes, err := client.ListStudents(ctx, &pb.Empty{})
	if err != nil {
		log.Fatalf("Error calling ListStudents: %v", err)
	}

	// Convert to JSON format
	data := map[string]interface{}{
		"students": listRes.Student,
	}

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		log.Fatalf("Error converting to JSON: %v", err)
	}

	log.Println(string(jsonData))
}