package services

import (
	"context"
	"sync"
	"time"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/university-library/grpc-service/pb"
)

type LoanServer struct {
	pb.UnimplementedLoanServiceServer
	mu    sync.RWMutex
	loans map[string]*pb.Loan
}

func NewLoanServer() *LoanServer {
	server := &LoanServer{
		loans: make(map[string]*pb.Loan),
	}
	
	// Add sample loans
	now := time.Now()
	sampleLoans := []*pb.Loan{
		{
			Id:         uuid.New().String(),
			StudentId:  "student-1",
			BookId:     "book-1",
			LoanDate:   now.AddDate(0, 0, -10).Format("2006-01-02"),
			ReturnDate: "",
			Status:     pb.LoanStatus_ONGOING,
		},
		{
			Id:         uuid.New().String(),
			StudentId:  "student-2",
			BookId:     "book-2",
			LoanDate:   now.AddDate(0, 0, -5).Format("2006-01-02"),
			ReturnDate: now.AddDate(0, 0, -1).Format("2006-01-02"),
			Status:     pb.LoanStatus_RETURNED,
		},
		{
			Id:         uuid.New().String(),
			StudentId:  "student-1",
			BookId:     "book-3",
			LoanDate:   now.AddDate(0, 0, -20).Format("2006-01-02"),
			ReturnDate: "",
			Status:     pb.LoanStatus_LATE,
		},
	}
	
	for _, loan := range sampleLoans {
		server.loans[loan.Id] = loan
	}
	
	return server
}

func (s *LoanServer) ListLoans(ctx context.Context, req *pb.ListLoansRequest) (*pb.ListLoansResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var loans []*pb.Loan
	for _, loan := range s.loans {
		// Filter by student ID if provided
		if req.StudentId != "" && loan.StudentId != req.StudentId {
			continue
		}
		loans = append(loans, loan)
	}

	return &pb.ListLoansResponse{
		Loans: loans,
	}, nil
}

func (s *LoanServer) GetLoan(ctx context.Context, req *pb.GetLoanRequest) (*pb.GetLoanResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	loan, exists := s.loans[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Loan with ID %s not found", req.Id)
	}

	return &pb.GetLoanResponse{
		Loan: loan,
	}, nil
}

func (s *LoanServer) CreateLoan(ctx context.Context, req *pb.CreateLoanRequest) (*pb.CreateLoanResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	// Validate required fields
	if req.StudentId == "" || req.BookId == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Student ID and Book ID are required")
	}

	// Create new loan
	loan := &pb.Loan{
		Id:         uuid.New().String(),
		StudentId:  req.StudentId,
		BookId:     req.BookId,
		LoanDate:   time.Now().Format("2006-01-02"),
		ReturnDate: "",
		Status:     pb.LoanStatus_ONGOING,
	}

	s.loans[loan.Id] = loan

	return &pb.CreateLoanResponse{
		Loan: loan,
	}, nil
}

func (s *LoanServer) ReturnLoan(ctx context.Context, req *pb.ReturnLoanRequest) (*pb.ReturnLoanResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	loan, exists := s.loans[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Loan with ID %s not found", req.Id)
	}

	if loan.Status == pb.LoanStatus_RETURNED {
		return nil, status.Errorf(codes.InvalidArgument, "Loan is already returned")
	}

	// Update loan status
	loan.ReturnDate = time.Now().Format("2006-01-02")
	loan.Status = pb.LoanStatus_RETURNED

	s.loans[req.Id] = loan

	return &pb.ReturnLoanResponse{
		Loan: loan,
	}, nil
} 