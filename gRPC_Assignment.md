# Protocol Buffers & gRPC Uygulama Geliştirme Ödevi

## 🎓 Ders: Açık Kaynak Kodlu Yazılımlar  
**Teslim Tarihi**: [Güncel teslim tarihi yazılacak]  
**Teslim Yeri**: Google Classroom (Yalnızca `DELIVERY.md` dosyası)

---

## 🎯 Ödevin Amacı

Bu ödevin amacı, öğrencilerin Protocol Buffers (Protobuf) dilini kullanarak bir API tanımı yapmaları ve bu tanıma uygun olarak gRPC protokolü ile çalışan bir istemci-sunucu uygulaması geliştirme becerilerini kazanmalarını sağlamaktır. Öğrenciler bir `.proto` dosyası oluşturacak ve bu tanıma uygun çalışan bir gRPC sunucu ve istemci uygulaması geliştireceklerdir.

---

## 📌 Ödev Senaryosu (Zorunlu Yapı)

Bir üniversiteye ait çevrim içi kütüphane sistemi için gRPC tabanlı bir servis geliştirilecektir. Aşağıdaki üç varlık ve operasyonlar zorunludur.

### 📚 Varlıklar ve Zorunlu Alanlar

#### 1. books
- `id` (string, UUID formatında)
- `title` (string)
- `author` (string)
- `isbn` (string, format: ISBN-13)
- `publisher` (string)
- `pageCount` (integer)
- `stock` (integer)

#### 2. students
- `id` (string, UUID formatında)
- `name` (string)
- `studentNumber` (string)
- `email` (string, format: email)
- `isActive` (boolean)

#### 3. loans
- `id` (string, UUID formatında)
- `studentId` (string)
- `bookId` (string)
- `loanDate` (string, format: date)
- `returnDate` (string, format: date, nullable)
- `status` (enum: ["ongoing", "returned", "late"])

### 🔗 Zorunlu Servis Metotları

Her varlık için en az aşağıdaki işlevleri sağlayan servis metotları `.proto` dosyasında tanımlanmalıdır:

#### 📘 books
- Listeleme
- Tekil görüntüleme
- Ekleme
- Güncelleme
- Silme

#### 👤 students
- Listeleme
- Tekil görüntüleme
- Ekleme
- Güncelleme
- Silme

#### 🔄 loans
- Listeleme
- Tekil görüntüleme
- Ödünç alma
- İade etme

Metot isimleri Türkçe veya İngilizce olabilir, tercihen İngilizce tercih edilmelidir.

---

## ✅ `.proto` Dosyanızda Bulunması Gerekenler

| Özellik                      | Gereklilik                         |
|-----------------------------|-------------------------------------|
| `syntax`, `package`, `option` tanımları | Zorunlu                   |
| En az 3 `service` tanımı    | books, students, loans              |
| CRUD ve işlevsel `rpc` tanımları | Yukarıdaki işlevleri kapsamalı |
| `message` tanımları         | Request ve response için            |
| `enum` kullanımı            | loans → `status` alanı              |
| Kod stili                   | Okunabilir, İngilizce tercih edilmeli |

---

## ✅ Uygulama Gereksinimleri

| Gereklilik                               | Açıklama                                                                  |
|------------------------------------------|---------------------------------------------------------------------------|
| Sunucu uygulaması                        | gRPC protokolüne uygun, mock veri ile çalışabilir                        |
| İstemci uygulaması                       | Servis metotlarını çağırmalı                                              |
| grpcurl ile test                         | Her servis `grpcurl` aracıyla test edilmeli                              |
| grpcurl test çıktılarının belgelenmesi  | Komutlar + ekran görüntüleri `grpcurl-tests.md` içinde sunulmalı         |
| Stub dosyaları repoya eklenmemeli        | Sadece `.proto` dosyası olmalı, stub kodlar build sırasında üretilmeli   |

> 💡 Not: Öğrenciler sunucu ve istemci uygulamalarını istedikleri programlama dilini kullanarak geliştirebilirler. (Java, Go, Python, Node.js, Rust vb.)

---

## 🗂️ GitHub Reposunda Bulunması Gereken Dosyalar

```
/
├── university.proto             # veya birden fazla .proto dosyası
├── README.md                   # Proje açıklaması
├── grpcurl-tests.md            # Test dokümantasyonu
├── src/
│   ├── server/
│   └── client/
└── DELIVERY.md                 # Sadece bu dosya Classroom’a yüklenecek
```

---

## 📝 Teslimat Formatı

Her öğrenci yalnızca `DELIVERY.md` adlı markdown dosyasını doldurarak Google Classroom’a yükleyecektir.  
Bu dosyada:  
- GitHub repo linki  
- grpcurl test dökümanı bilgisi  
- Projeyi çalıştırmak için yönergeler  
yer almalıdır.

---

## 🧮 Değerlendirme Kriterleri

| Kriter                                         | Puan |
|-----------------------------------------------|------|
| `.proto` tanımının doğruluğu ve kapsamı       | 30   |
| Sunucu uygulamasının çalışabilirliği          | 20   |
| İstemci uygulamasının fonksiyonelliği         | 20   |
| grpcurl test belgesi (`grpcurl-tests.md`)     | 20   |
| Repo yapısı, kod düzeni ve okunabilirlik      | 10   |
| **Toplam**                                    | **100** |

---

Başarılar!
