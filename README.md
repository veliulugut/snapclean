# SnapClean - Veri Temizleme Aracı

**Dil Seçimi | Language Selection**

[Türkçe](#turkish) | [English](#english)

---

## Türkçe Dokümantasyon {#turkish}

### Proje Açıklaması

SnapClean, CSV ve Excel dosyalarınızda bulunan veri temizleme işlemlerini interaktif ve kullanıcı dostu bir arayüz üzerinden gerçekleştirmenizi sağlayan Go dilinde yazılmış bir masaüstü uygulamasıdır.

### Ana Özellikler

- **Çoklu Format Desteği**: CSV ve Excel (.xlsx, .xls) dosyalarını destekler
- **Etkileşimli Kullanıcı Arayüzü**: Terminal tabanlı modern ve yanıtlı arayüz (Bubble Tea)
- **Kapsamlı Temizleme İşlemleri**:
  - Boş satırların kaldırılması
  - Boş sütunların kaldırılması
  - Veri tutarlılığı kontrolü
  - Sütun seçimi ve yönetimi
- **Dosya Yönetimi**: Grafik dosya seçici ile kolayca dosya açma
- **Veri Görselleştirme**: Temizlenen verileri tablo formatında görüntüleme

### Gereksinimler

- Go 1.24.1 veya daha yeni sürüm
- macOS, Linux veya Windows işletim sistemi

### Kurulum

#### 1. Depoyu Klonlayın

```bash
git clone https://github.com/veliulugut/snapclean.git
cd snapclean
```

#### 2. Bağımlılıkları Yükleyin

```bash
go mod download
go mod tidy
```

#### 3. Uygulamayı Çalıştırın

```bash
go run cmd/main.go
```

#### 4. (Opsiyonel) Yürütülebilir Dosya Oluşturun

```bash
go build -o snapclean cmd/main.go
./snapclean
```

### Kullanım

1. **Uygulamayı Başlatın**: `go run cmd/main.go` komutunu çalıştırın
2. **Dosya Seçin**: Splash ekranından dosya seçici ile CSV veya Excel dosyası açın
3. **Temizleme İşlemleri**: Ana menüden istediğiniz temizleme işlemlerini seçin
4. **Sonuçları Görüntüleyin**: Tablo görünümünde temizlenmiş verilerinizi kontrol edin
5. **Dosyayı Kaydedin**: Temizlenmiş dosyayı istediğiniz formatta dışarı aktarın

### Proje Yapısı

```
snapclean/
├── cmd/
│   └── main.go                 # Uygulamanın giriş noktası
├── internal/
│   ├── cleaner/                # Veri temizleme işlevleri
│   ├── file/                   # Dosya yükleme ve kaydetme
│   ├── models/                 # Veri yapıları
│   ├── reshaper/               # Veri şekillendirme işlemleri
│   ├── summarizer/             # Veri özeti oluşturma
│   └── tui/                    # Terminal kullanıcı arayüzü bileşenleri
├── media/                      # Proje görselleri
├── go.mod                      # Go modülü tanımı
└── README.md                   # Bu dosya
```

### Teknoloji Stack

- **Go 1.24.1**: Programlama dili
- **Bubble Tea**: Terminal UI framework
- **Lipgloss**: Terminal stil ve renk kütüphanesi
- **Excelize**: Excel dosya işleme
- **Zenity**: Grafik dosya seçici dialog

### Geliştirme

#### Testleri Çalıştırın

```bash
go test ./...
go test -v ./...  # Detaylı çıktı
```

#### Kodu Biçimlendirin

```bash
go fmt ./...
```

#### Kodu Analiz Edin

```bash
go vet ./...
```

### Lisans

Bu proje açık kaynak olarak sunulmaktadır.

### Katkıda Bulunun

Katkılarınız hoşlanır! Lütfen:
1. Depoyu fork edin
2. Özellik dalı oluşturun (`git checkout -b feature/YeniÖzellik`)
3. Değişikliklerinizi commit edin (`git commit -am 'Yeni özellik ekle'`)
4. Dala push yapın (`git push origin feature/YeniÖzellik`)
5. Pull Request oluşturun

### İletişim

Sorularınız veya önerileriniz için bir issue oluşturun.

---

## English Documentation {#english}

### Project Description

SnapClean is a desktop application written in Go that allows you to perform data cleaning operations on your CSV and Excel files through an interactive and user-friendly interface.

### Key Features

- **Multiple Format Support**: Supports CSV and Excel (.xlsx, .xls) files
- **Interactive User Interface**: Modern and responsive terminal-based UI (Bubble Tea)
- **Comprehensive Cleaning Operations**:
  - Remove empty rows
  - Remove empty columns
  - Data consistency checking
  - Column selection and management
- **File Management**: Easy file opening with graphical file picker
- **Data Visualization**: View cleaned data in table format

### Requirements

- Go 1.24.1 or later
- macOS, Linux, or Windows operating system

### Installation

#### 1. Clone the Repository

```bash
git clone https://github.com/veliulugut/snapclean.git
cd snapclean
```

#### 2. Download Dependencies

```bash
go mod download
go mod tidy
```

#### 3. Run the Application

```bash
go run cmd/main.go
```

#### 4. (Optional) Build an Executable

```bash
go build -o snapclean cmd/main.go
./snapclean
```

### Usage

1. **Start the Application**: Run the command `go run cmd/main.go`
2. **Select a File**: Open a CSV or Excel file using the file picker from the splash screen
3. **Cleaning Operations**: Choose the cleaning operations you want from the main menu
4. **View Results**: Check your cleaned data in the table view
5. **Save the File**: Export the cleaned file in your desired format

### Project Structure

```
snapclean/
├── cmd/
│   └── main.go                 # Application entry point
├── internal/
│   ├── cleaner/                # Data cleaning functions
│   ├── file/                   # File loading and saving
│   ├── models/                 # Data structures
│   ├── reshaper/               # Data reshaping operations
│   ├── summarizer/             # Data summary creation
│   └── tui/                    # Terminal UI components
├── media/                      # Project images
├── go.mod                      # Go module definition
└── README.md                   # This file
```

### Technology Stack

- **Go 1.24.1**: Programming language
- **Bubble Tea**: Terminal UI framework
- **Lipgloss**: Terminal styling and color library
- **Excelize**: Excel file processing
- **Zenity**: Graphical file picker dialog

### Development

#### Run Tests

```bash
go test ./...
go test -v ./...  # Verbose output
```

#### Format Code

```bash
go fmt ./...
```

#### Analyze Code

```bash
go vet ./...
```

### License

This project is provided as open source.

### Contributing

Contributions are welcome! Please:
1. Fork the repository
2. Create a feature branch (`git checkout -b feature/NewFeature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/NewFeature`)
5. Create a Pull Request

### Contact

Please create an issue for any questions or suggestions.
