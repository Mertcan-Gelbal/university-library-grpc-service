package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/university-library/grpc-service/pb"
)

const (
	address = "localhost:50051"
)

func main() {
	// Set up a connection to the server
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	// Create service clients
	bookClient := pb.NewBookServiceClient(conn)
	studentClient := pb.NewStudentServiceClient(conn)
	loanClient := pb.NewLoanServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	// Test Book Service
	log.Println("=== Testing Book Service ===")
	testBookService(ctx, bookClient)

	// Test Student Service
	log.Println("\n=== Testing Student Service ===")
	testStudentService(ctx, studentClient)

	// Test Loan Service
	log.Println("\n=== Testing Loan Service ===")
	testLoanService(ctx, loanClient)
}

func testBookService(ctx context.Context, client pb.BookServiceClient) {
	// List books
	log.Println("Listing books...")
	listResp, err := client.ListBooks(ctx, &pb.ListBooksRequest{})
	if err != nil {
		log.Printf("Error listing books: %v", err)
		return
	}
	log.Printf("Found %d books", len(listResp.Books))
	for _, book := range listResp.Books {
		log.Printf("- %s by %s (Stock: %d)", book.Title, book.Author, book.Stock)
	}

	// Create a new book
	log.Println("\nCreating a new book...")
	newBook := &pb.Book{
		Title:     "Effective Go",
		Author:    "The Go Team",
		Isbn:      "978-0000000000",
		Publisher: "Google",
		PageCount: 250,
		Stock:     10,
	}
	createResp, err := client.CreateBook(ctx, &pb.CreateBookRequest{Book: newBook})
	if err != nil {
		log.Printf("Error creating book: %v", err)
		return
	}
	log.Printf("Created book with ID: %s", createResp.Book.Id)

	// Get the created book
	log.Println("Getting the created book...")
	getResp, err := client.GetBook(ctx, &pb.GetBookRequest{Id: createResp.Book.Id})
	if err != nil {
		log.Printf("Error getting book: %v", err)
		return
	}
	log.Printf("Retrieved book: %s", getResp.Book.Title)
}

func testStudentService(ctx context.Context, client pb.StudentServiceClient) {
	// List students
	log.Println("Listing students...")
	listResp, err := client.ListStudents(ctx, &pb.ListStudentsRequest{})
	if err != nil {
		log.Printf("Error listing students: %v", err)
		return
	}
	log.Printf("Found %d students", len(listResp.Students))
	for _, student := range listResp.Students {
		log.Printf("- %s (%s) - Active: %t", student.Name, student.StudentNumber, student.IsActive)
	}

	// Create a new student
	log.Println("\nCreating a new student...")
	newStudent := &pb.Student{
		Name:          "Fatma Özkan",
		StudentNumber: "20230001",
		Email:         "fatma.ozkan@university.edu.tr",
		IsActive:      true,
	}
	createResp, err := client.CreateStudent(ctx, &pb.CreateStudentRequest{Student: newStudent})
	if err != nil {
		log.Printf("Error creating student: %v", err)
		return
	}
	log.Printf("Created student with ID: %s", createResp.Student.Id)
}

func testLoanService(ctx context.Context, client pb.LoanServiceClient) {
	// List loans
	log.Println("Listing loans...")
	listResp, err := client.ListLoans(ctx, &pb.ListLoansRequest{})
	if err != nil {
		log.Printf("Error listing loans: %v", err)
		return
	}
	log.Printf("Found %d loans", len(listResp.Loans))
	for _, loan := range listResp.Loans {
		log.Printf("- Loan ID: %s, Student: %s, Book: %s, Status: %s", 
			loan.Id, loan.StudentId, loan.BookId, loan.Status.String())
	}

	// Create a new loan
	log.Println("\nCreating a new loan...")
	createResp, err := client.CreateLoan(ctx, &pb.CreateLoanRequest{
		StudentId: "student-test",
		BookId:    "book-test",
	})
	if err != nil {
		log.Printf("Error creating loan: %v", err)
		return
	}
	log.Printf("Created loan with ID: %s", createResp.Loan.Id)

	// Return the loan
	log.Println("Returning the loan...")
	returnResp, err := client.ReturnLoan(ctx, &pb.ReturnLoanRequest{Id: createResp.Loan.Id})
	if err != nil {
		log.Printf("Error returning loan: %v", err)
		return
	}
	log.Printf("Loan returned successfully. Return date: %s", returnResp.Loan.ReturnDate)
} 