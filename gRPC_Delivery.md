# gRPC Uygulama Geliştirme Ödevi Teslim Raporu

## 👤 Öğrenci Bilgileri
- **Ad Soyad**: Mertcan Gelbal
- **Öğrenci Numarası**: 171421012
- **Kullanılan Programlama Dili**: Go (Golang)

---

## 📦 GitHub Repo

Proje tamamı GitHub reposuna yüklenmiştir. Stub kodlar hariç tutulmuş, sadece kaynak kodlar ve .proto dosyası dahil edilmiştir.

### 🔗 GitHub Repo Linki
https://github.com/Mertcan-Gelbal/university-library-grpc-service

---

## 📄 .proto Dosyası

- **.proto dosyasının adı**: `university.proto`
- **Tanımlanan servisler ve metod sayısı**: 
  - BookService (5 metod): ListBooks, GetBook, CreateBook, UpdateBook, DeleteBook
  - StudentService (5 metod): ListStudents, GetStudent, CreateStudent, UpdateStudent, DeleteStudent
  - LoanService (4 metod): ListLoans, GetLoan, CreateLoan, ReturnLoan
  - **Toplam**: 3 servis, 14 metod
- **Enum kullanımı**: Evet, `LoanStatus` enum'u `Loan` mesajında kullanılmıştır (ONGOING, RETURNED, LATE)
- **Dil kullanımı**: İngilizce tercih edilmiştir. Tüm servis isimleri, metod isimleri ve alan isimleri İngilizce olarak tanımlanmıştır.

---

## 🧪 grpcurl Test Dokümantasyonu

Aşağıdaki bilgiler `grpcurl-tests.md` adlı ayrı bir markdown dosyasında detaylı olarak yer almaktadır:

- Her metot için kullanılan `grpcurl` komutu
- Dönen yanıtların örnek çıktıları
- Hatalı durum senaryoları (404, boş yanıt, validation hataları vb.)
- Enum değerlerinin test edilmesi
- Filtreleme özelliklerinin test edilmesi

Test dosyası kapsamlı şekilde hazırlanmış olup, tüm CRUD operasyonları ve hata senaryoları dahil edilmiştir.

---

## 🛠️ Derleme ve Çalıştırma Adımları

Projeyi `.proto` dosyasından derleyip sunucu/istemci uygulamasını çalıştırmak için gereken komutlar:

```bash
# 1. Bağımlılıkları yükle
go mod tidy

# 2. Protocol Buffers kodlarını oluştur
protoc --go_out=. --go-grpc_out=. university.proto

# 3. Sunucuyu çalıştır (Terminal 1)
go run src/server/main.go

# 4. İstemciyi test et (Terminal 2)
go run src/client/main.go

# 5. grpcurl ile test et
grpcurl -plaintext localhost:50051 list
grpcurl -plaintext localhost:50051 university.library.BookService/ListBooks
```

**Not**: Sunucu `localhost:50051` portunda çalışmaktadır ve reflection servisi aktif edilmiştir.

---

## ⚠️ Kontrol Listesi

- [x] Stub dosyaları GitHub reposuna eklenmedi  
- [x] grpcurl komutları test belgesinde yer alıyor  
- [x] Örnek çıktılar test belgesine eklendi  
- [x] Tüm servisler çalışır durumda  
- [x] README.md içinde yeterli açıklama var  
- [x] Enum kullanımı implementasyonu yapılmış
- [x] Validation kontrolleri eklendi
- [x] Error handling gRPC standartlarına uygun
- [x] Thread safety sağlandı (mutex kullanımı)
- [x] Clean code prensipleri uygulandı

---

## 📌 Ek Açıklamalar

### Teknik Kararlar

1. **Programlama Dili**: Go seçilmiştir çünkü gRPC ile mükemmel entegrasyona sahiptir ve performans açısından avantajlıdır.

2. **Veri Yönetimi**: Gerçek bir veritabanı yerine in-memory veri yapıları kullanılmıştır. Bu yaklaşım prototip geliştirme için uygundur ve bağımlılıkları minimize eder.

3. **Thread Safety**: Concurrent erişim için `sync.RWMutex` kullanılarak thread safety sağlanmıştır.

4. **Error Handling**: gRPC status kodları (NotFound, InvalidArgument) kullanılarak standart hata yönetimi implementasyonu yapılmıştır.

5. **Validation**: Tüm servislerde input validation kontrolleri eklenmiştir.

### Mimari Özellikler

- **Clean Architecture**: Servis katmanları ayrılmış, bağımlılıklar minimize edilmiştir
- **Constructor Pattern**: Servis instance'ları için constructor pattern kullanılmıştır
- **Reflection Support**: grpcurl desteği için reflection servisi etkinleştirilmiştir
- **Pagination Ready**: Temel pagination desteği eklenmiştir (gelecek geliştirmeler için)

### Test Coverage

- Tüm CRUD operasyonları test edilmiştir
- Hata senaryoları kapsamlı şekilde test edilmiştir
- Enum değerleri doğru çalışmaktadır
- Filtreleme özellikleri test edilmiştir

Proje, ödev gereksinimlerinin tamamını karşılamakta ve production-ready kod kalitesinde geliştirilmiştir.

---

Teşekkürler!
