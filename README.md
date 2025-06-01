# University Library gRPC Service

This project is a microservice application developed using Protocol Buffers and gRPC technologies for a university library system.

## Project Structure

```
/
├── university.proto         # Protocol Buffers definition file
├── go.mod                  # Go module definitions
├── README.md               # Project documentation
├── grpcurl-tests.md        # grpcurl test documentation
├── src/
│   ├── server/
│   │   ├── main.go         # gRPC server application
│   │   └── services/       # Service implementations
│   │       ├── book_service.go
│   │       ├── student_service.go
│   │       └── loan_service.go
│   └── client/
│       └── main.go         # gRPC client application
└── DELIVERY.md             # Assignment delivery report
```

## Features

### Entities

1. **Books**
   - ID (UUID)
   - Title
   - Author
   - ISBN
   - Publisher
   - Page count
   - Stock quantity

2. **Students**
   - ID (UUID)
   - Name
   - Student number
   - Email
   - Active status

3. **Loans**
   - ID (UUID)
   - Student ID
   - Book ID
   - Loan date
   - Return date
   - Status (ongoing, returned, late)

### Services

#### BookService
- `ListBooks` - List books
- `GetBook` - Get single book
- `CreateBook` - Add new book
- `UpdateBook` - Update book
- `DeleteBook` - Delete book

#### StudentService
- `ListStudents` - List students
- `GetStudent` - Get single student
- `CreateStudent` - Add new student
- `UpdateStudent` - Update student
- `DeleteStudent` - Delete student

#### LoanService
- `ListLoans` - List loan transactions
- `GetLoan` - Get single loan transaction
- `CreateLoan` - Create new loan transaction
- `ReturnLoan` - Return book

## Requirements

- Go 1.21 or higher
- Protocol Buffers compiler (protoc)
- grpcurl (for testing)

## Installation and Running

### 1. Install Dependencies

```bash
go mod tidy
```

### 2. Generate Protocol Buffers Code

```bash
protoc --go_out=. --go-grpc_out=. university.proto
```

### 3. Run Server

```bash
go run src/server/main.go
```

The server will start listening on `localhost:50051`.

### 4. Test Client

Open a new terminal and run:

```bash
go run src/client/main.go
```

## Testing with grpcurl

While the server is running, you can test services with grpcurl:

### List Services

```bash
grpcurl -plaintext localhost:50051 list
```

### List Books

```bash
grpcurl -plaintext localhost:50051 university.library.BookService/ListBooks
```

### Add New Book

```bash
grpcurl -plaintext -d '{
  "book": {
    "title": "Test Book",
    "author": "Test Author",
    "isbn": "978-0000000000",
    "publisher": "Test Publisher",
    "page_count": 200,
    "stock": 5
  }
}' localhost:50051 university.library.BookService/CreateBook
```

For detailed test commands, see `grpcurl-tests.md` file.

## Technology Stack

- **Go**: Main programming language
- **Protocol Buffers**: Data serialization
- **gRPC**: RPC framework
- **UUID**: Unique identifier generation

## Architecture Decisions

1. **Clean Architecture**: Service layers are separated, dependencies are minimized.
2. **Thread Safety**: Mutex is used for concurrent access.
3. **Error Handling**: Proper error management with gRPC status codes.
4. **Mock Data**: In-memory data structures are used instead of real database.
5. **Reflection**: Reflection service is enabled for grpcurl support.

## Development Notes

- Stub files are not included in the repository
- All services support CRUD operations
- Enum usage is implemented for loan status
- Basic pagination support is added
- Validation controls are implemented

## License

This project is developed for educational purposes. 