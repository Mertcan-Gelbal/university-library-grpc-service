# grpcurl Test Dokümantasyonu

Bu dokümanda University Library gRPC servisinin grpcurl ile test edilmesi için gerekli komutlar ve beklenen yanıtlar yer almaktadır.

## Ön Koşullar

1. gRPC sunucusunun çalışıyor olması (`go run src/server/main.go`)
2. grpcurl'ün sisteme kurulu olması

## Temel Komutlar

### Servisleri Listele

```bash
grpcurl -plaintext localhost:50051 list
```

**Beklenen Çıktı:**
```
grpc.reflection.v1alpha.ServerReflection
university.library.BookService
university.library.LoanService
university.library.StudentService
```

### Servis Metotlarını Listele

```bash
grpcurl -plaintext localhost:50051 list university.library.BookService
```

**Beklenen Çıktı:**
```
university.library.BookService.CreateBook
university.library.BookService.DeleteBook
university.library.BookService.GetBook
university.library.BookService.ListBooks
university.library.BookService.UpdateBook
```

## BookService Testleri

### 1. Kitapları Listele

```bash
grpcurl -plaintext localhost:50051 university.library.BookService/ListBooks
```

**Beklenen Çıktı:**
```json
{
  "books": [
    {
      "id": "uuid-1",
      "title": "Clean Code",
      "author": "Robert C. Martin",
      "isbn": "978-0132350884",
      "publisher": "Prentice Hall",
      "pageCount": 464,
      "stock": 5
    },
    {
      "id": "uuid-2",
      "title": "Design Patterns",
      "author": "Gang of Four",
      "isbn": "978-0201633610",
      "publisher": "Addison-Wesley",
      "pageCount": 395,
      "stock": 3
    },
    {
      "id": "uuid-3",
      "title": "The Go Programming Language",
      "author": "Alan Donovan, Brian Kernighan",
      "isbn": "978-0134190440",
      "publisher": "Addison-Wesley",
      "pageCount": 380,
      "stock": 7
    }
  ]
}
```

### 2. Yeni Kitap Oluştur

```bash
grpcurl -plaintext -d '{
  "book": {
    "title": "Effective Go",
    "author": "The Go Team",
    "isbn": "978-0000000000",
    "publisher": "Google",
    "page_count": 250,
    "stock": 10
  }
}' localhost:50051 university.library.BookService/CreateBook
```

**Beklenen Çıktı:**
```json
{
  "book": {
    "id": "generated-uuid",
    "title": "Effective Go",
    "author": "The Go Team",
    "isbn": "978-0000000000",
    "publisher": "Google",
    "pageCount": 250,
    "stock": 10
  }
}
```

### 3. Kitap Getir (ID ile)

```bash
grpcurl -plaintext -d '{
  "id": "BOOK_ID_BURAYA"
}' localhost:50051 university.library.BookService/GetBook
```

### 4. Kitap Güncelle

```bash
grpcurl -plaintext -d '{
  "book": {
    "id": "BOOK_ID_BURAYA",
    "title": "Updated Title",
    "author": "Updated Author",
    "isbn": "978-1111111111",
    "publisher": "Updated Publisher",
    "page_count": 300,
    "stock": 15
  }
}' localhost:50051 university.library.BookService/UpdateBook
```

### 5. Kitap Sil

```bash
grpcurl -plaintext -d '{
  "id": "BOOK_ID_BURAYA"
}' localhost:50051 university.library.BookService/DeleteBook
```

**Beklenen Çıktı:**
```json
{
  "success": true,
  "message": "Book deleted successfully"
}
```

## StudentService Testleri

### 1. Öğrencileri Listele

```bash
grpcurl -plaintext localhost:50051 university.library.StudentService/ListStudents
```

**Beklenen Çıktı:**
```json
{
  "students": [
    {
      "id": "uuid-1",
      "name": "Ahmet Yılmaz",
      "studentNumber": "20210001",
      "email": "ahmet.yilmaz@university.edu.tr",
      "isActive": true
    },
    {
      "id": "uuid-2",
      "name": "Ayşe Demir",
      "studentNumber": "20210002",
      "email": "ayse.demir@university.edu.tr",
      "isActive": true
    },
    {
      "id": "uuid-3",
      "name": "Mehmet Kaya",
      "studentNumber": "20200015",
      "email": "mehmet.kaya@university.edu.tr",
      "isActive": false
    }
  ]
}
```

### 2. Yeni Öğrenci Oluştur

```bash
grpcurl -plaintext -d '{
  "student": {
    "name": "Fatma Özkan",
    "student_number": "20230001",
    "email": "fatma.ozkan@university.edu.tr",
    "is_active": true
  }
}' localhost:50051 university.library.StudentService/CreateStudent
```

### 3. Öğrenci Getir

```bash
grpcurl -plaintext -d '{
  "id": "STUDENT_ID_BURAYA"
}' localhost:50051 university.library.StudentService/GetStudent
```

### 4. Öğrenci Güncelle

```bash
grpcurl -plaintext -d '{
  "student": {
    "id": "STUDENT_ID_BURAYA",
    "name": "Updated Name",
    "student_number": "20230002",
    "email": "updated@university.edu.tr",
    "is_active": false
  }
}' localhost:50051 university.library.StudentService/UpdateStudent
```

### 5. Öğrenci Sil

```bash
grpcurl -plaintext -d '{
  "id": "STUDENT_ID_BURAYA"
}' localhost:50051 university.library.StudentService/DeleteStudent
```

## LoanService Testleri

### 1. Ödünç İşlemlerini Listele

```bash
grpcurl -plaintext localhost:50051 university.library.LoanService/ListLoans
```

**Beklenen Çıktı:**
```json
{
  "loans": [
    {
      "id": "uuid-1",
      "studentId": "student-1",
      "bookId": "book-1",
      "loanDate": "2024-01-01",
      "returnDate": "",
      "status": "ONGOING"
    },
    {
      "id": "uuid-2",
      "studentId": "student-2",
      "bookId": "book-2",
      "loanDate": "2024-01-05",
      "returnDate": "2024-01-10",
      "status": "RETURNED"
    },
    {
      "id": "uuid-3",
      "studentId": "student-1",
      "bookId": "book-3",
      "loanDate": "2023-12-15",
      "returnDate": "",
      "status": "LATE"
    }
  ]
}
```

### 2. Yeni Ödünç İşlemi Oluştur

```bash
grpcurl -plaintext -d '{
  "student_id": "student-test",
  "book_id": "book-test"
}' localhost:50051 university.library.LoanService/CreateLoan
```

**Beklenen Çıktı:**
```json
{
  "loan": {
    "id": "generated-uuid",
    "studentId": "student-test",
    "bookId": "book-test",
    "loanDate": "2024-01-15",
    "returnDate": "",
    "status": "ONGOING"
  }
}
```

### 3. Ödünç İşlemi Getir

```bash
grpcurl -plaintext -d '{
  "id": "LOAN_ID_BURAYA"
}' localhost:50051 university.library.LoanService/GetLoan
```

### 4. Kitap İade Et

```bash
grpcurl -plaintext -d '{
  "id": "LOAN_ID_BURAYA"
}' localhost:50051 university.library.LoanService/ReturnLoan
```

**Beklenen Çıktı:**
```json
{
  "loan": {
    "id": "loan-uuid",
    "studentId": "student-test",
    "bookId": "book-test",
    "loanDate": "2024-01-15",
    "returnDate": "2024-01-16",
    "status": "RETURNED"
  }
}
```

### 5. Öğrenciye Göre Ödünç İşlemlerini Filtrele

```bash
grpcurl -plaintext -d '{
  "student_id": "student-1"
}' localhost:50051 university.library.LoanService/ListLoans
```

## Hata Senaryoları

### 1. Olmayan Kitap Getirme

```bash
grpcurl -plaintext -d '{
  "id": "nonexistent-id"
}' localhost:50051 university.library.BookService/GetBook
```

**Beklenen Hata:**
```
ERROR:
  Code: NotFound
  Message: Book with ID nonexistent-id not found
```

### 2. Geçersiz Veri ile Kitap Oluşturma

```bash
grpcurl -plaintext -d '{
  "book": {
    "title": "",
    "author": "",
    "isbn": ""
  }
}' localhost:50051 university.library.BookService/CreateBook
```

**Beklenen Hata:**
```
ERROR:
  Code: InvalidArgument
  Message: Title, author, and ISBN are required
```

### 3. Zaten İade Edilmiş Kitabı İade Etme

```bash
grpcurl -plaintext -d '{
  "id": "RETURNED_LOAN_ID"
}' localhost:50051 university.library.LoanService/ReturnLoan
```

**Beklenen Hata:**
```
ERROR:
  Code: InvalidArgument
  Message: Loan is already returned
```

## Test Sonuçları

Tüm testler başarıyla çalıştırıldığında:
- ✅ Tüm servisler erişilebilir
- ✅ CRUD operasyonları çalışıyor
- ✅ Validation kontrolleri aktif
- ✅ Hata yönetimi uygun
- ✅ Enum değerleri doğru çalışıyor
- ✅ Reflection servisi aktif 