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

type BookServer struct {
	pb.UnimplementedBookServiceServer
	mu    sync.RWMutex
	books map[string]*pb.Book
}

func NewBookServer() *BookServer {
	server := &BookServer{
		books: make(map[string]*pb.Book),
	}
	
	// Add sample books
	sampleBooks := []*pb.Book{
		{
			Id:        uuid.New().String(),
			Title:     "Clean Code",
			Author:    "Robert C. Martin",
			Isbn:      "978-0132350884",
			Publisher: "Prentice Hall",
			PageCount: 464,
			Stock:     5,
		},
		{
			Id:        uuid.New().String(),
			Title:     "Design Patterns",
			Author:    "Gang of Four",
			Isbn:      "978-0201633610",
			Publisher: "Addison-Wesley",
			PageCount: 395,
			Stock:     3,
		},
		{
			Id:        uuid.New().String(),
			Title:     "The Go Programming Language",
			Author:    "Alan Donovan, Brian Kernighan",
			Isbn:      "978-0134190440",
			Publisher: "Addison-Wesley",
			PageCount: 380,
			Stock:     7,
		},
	}
	
	for _, book := range sampleBooks {
		server.books[book.Id] = book
	}
	
	return server
}

func (s *BookServer) ListBooks(ctx context.Context, req *pb.ListBooksRequest) (*pb.ListBooksResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	var books []*pb.Book
	for _, book := range s.books {
		books = append(books, book)
	}

	return &pb.ListBooksResponse{
		Books: books,
	}, nil
}

func (s *BookServer) GetBook(ctx context.Context, req *pb.GetBookRequest) (*pb.GetBookResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	book, exists := s.books[req.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Book with ID %s not found", req.Id)
	}

	return &pb.GetBookResponse{
		Book: book,
	}, nil
}

func (s *BookServer) CreateBook(ctx context.Context, req *pb.CreateBookRequest) (*pb.CreateBookResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Book == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Book data is required")
	}

	// Generate new ID if not provided
	if req.Book.Id == "" {
		req.Book.Id = uuid.New().String()
	}

	// Validate required fields
	if req.Book.Title == "" || req.Book.Author == "" || req.Book.Isbn == "" {
		return nil, status.Errorf(codes.InvalidArgument, "Title, author, and ISBN are required")
	}

	s.books[req.Book.Id] = req.Book

	return &pb.CreateBookResponse{
		Book: req.Book,
	}, nil
}

func (s *BookServer) UpdateBook(ctx context.Context, req *pb.UpdateBookRequest) (*pb.UpdateBookResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if req.Book == nil {
		return nil, status.Errorf(codes.InvalidArgument, "Book data is required")
	}

	_, exists := s.books[req.Book.Id]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "Book with ID %s not found", req.Book.Id)
	}

	s.books[req.Book.Id] = req.Book

	return &pb.UpdateBookResponse{
		Book: req.Book,
	}, nil
}

func (s *BookServer) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*pb.DeleteBookResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.books[req.Id]
	if !exists {
		return &pb.DeleteBookResponse{
			Success: false,
			Message: fmt.Sprintf("Book with ID %s not found", req.Id),
		}, nil
	}

	delete(s.books, req.Id)

	return &pb.DeleteBookResponse{
		Success: true,
		Message: "Book deleted successfully",
	}, nil
} 