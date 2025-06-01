package services

import (
	"context"
	"fmt"
	"sync"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/university-library/grpc-service/pb"
)

type StudentServer struct {
	pb.UnimplementedStudentServiceServer
	mu       sync.RWMutex
	students map[string]*pb.Student
}

func NewStudentServer() *StudentServer {
	server := &StudentServer{
		students: make(map[string]*pb.Student),
	}
	
	// Add sample students
	sampleStudents := []*pb.Student{
		{
			Id:            uuid.New().String(),
			Name:          "Ahmet Yılmaz",
			StudentNumber: "20210001",
			Email:         "ahmet.yilmaz@university.edu.tr",
			IsActive:      true,
		},
		{
			Id:            uuid.New().String(),
			Name:          "Ayşe Demir",
			StudentNumber: "20210002",
			Email:         "ayse.demir@university.edu.tr",
			IsActive:      true,
		},
		{
			Id:            uuid.New().String(),
			Name:          "Mehmet Kaya",
			StudentNumber: "20200015",
			Email:         "mehmet.kaya@university.edu.tr",
			IsActive:      false,
		},
	}
	
	for _, student := range sampleStudents {
		server.students[student.Id] = student
	}
	
	return server
}

func (s *StudentServer) ListStudents(ctx context.Context, req *pb.ListStudentsRequest) (*pb.ListStudentsResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var students []*pb.Student
	for _, student := range s.students {
		students = append(students, student)
	}

	return &pb.ListStudentsResponse{
		Students: students,
	}, nil
}

func (s *StudentServer) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	student, exists := s.students[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Student with ID %s not found", req.Id)
	}

	return &pb.GetStudentResponse{
		Student: student,
	}, nil
}

func (s *StudentServer) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Student == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Student data is required")
	}

	// Generate new ID if not provided
	if req.Student.Id == "" {
		req.Student.Id = uuid.New().String()
	}

	// Validate required fields
	if req.Student.Name == "" || req.Student.StudentNumber == "" || req.Student.Email == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Name, student number, and email are required")
	}

	s.students[req.Student.Id] = req.Student

	return &pb.CreateStudentResponse{
		Student: req.Student,
	}, nil
}

func (s *StudentServer) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Student == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Student data is required")
	}

	_, exists := s.students[req.Student.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Student with ID %s not found", req.Student.Id)
	}

	s.students[req.Student.Id] = req.Student

	return &pb.UpdateStudentResponse{
		Student: req.Student,
	}, nil
}

func (s *StudentServer) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.students[req.Id]
	if !exists {
		return &pb.DeleteStudentResponse{
			Success: false,
			Message: fmt.Sprintf("Student with ID %s not found", req.Id),
		}, nil
	}

	delete(s.students, req.Id)

	return &pb.DeleteStudentResponse{
		Success: true,
		Message: "Student deleted successfully",
	}, nil
} 