# SnapClean
MVP İçeriği

Bubble Tea TUI

Ana menü

Dosya seçimi → GUI file picker ile (Windows/Mac/Linux)

Tablo görüntüleme, ok tuşları ile navigasyon

CSV/Excel yükleme ve kaydetme

CSV: encoding/csv

Excel: excelize

Veri Temizleme Fonksiyonları

Boş satır/sütun temizleme

Sütun isimlerini normalize etme

Duplicate satır kontrolü

Export

Temizlenmiş / özetlenmiş dosyayı CSV veya XLSX olarak kaydetme

Temel proje dizini ve modüler yapı

cmd/ → CLI komutları

internal/ → temizleme, utils ve file loader

examples/ → örnek dosyalar

Go Data Cleaner – Interactive CSV/Excel Tool
1️⃣ Proje Amacı

Veri analistlerinin ve veri mühendislerinin günlük Excel/CSV işlemlerinde yaşadığı veri temizleme, özetleme ve QA sorunlarını çözmek.

Terminal TUI üzerinden interaktif kullanım ile komut yazmak zorunda kalmadan işlemleri hızlıca yapmak.

Büyük dosyaları Golang’in hızını ve concurrency yeteneklerini kullanarak anlık olarak işlemek.

2️⃣ Hedef Kullanıcılar

Veri analistleri (Excel, CSV, Data Warehouse kullanıcıları)

ETL pipeline veya SQL veri testi yapan veri mühendisleri

Büyük veri dosyalarıyla çalışan ve hızlı, interaktif çözümler isteyen profesyoneller

3️⃣ Temel Özellikler (MVP)
A) CSV / Excel Yükleme

Kullanıcı “Load CSV/Excel” seçeneğine geldiğinde GUI file picker açılır.

Windows, Mac ve Linux üzerinde OS-native file dialog kullanılır.

Seçilen dosya path’i alınır ve terminal TUI’de tablo olarak görüntülenir.

B) Tablo Görüntüleme

Bubble Tea TUI ile interaktif tablo görünümü

Ok tuşları ile satır ve sütun navigasyonu

Seçili alan highlight edilir

C) Veri Temizleme

Boş satır ve sütunları silme

Sütun isimlerini normalize etme (küçük harf, alt çizgi, özel karakter temizleme)

Yinelenen satırları silme

D) Veri Testleri / QA

Duplicate kontrol

Missing values kontrolü

Data multiplication check (özellikle data warehouse testleri için)

E) Özetleme / Pivot

Tek satır komut yerine interaktif menü üzerinden özet tablolar oluşturma

Unique / total / no-show / cancelled gibi metrikleri hesaplama

F) Export

Temizlenmiş ve özetlenmiş veriyi CSV veya Excel olarak kaydetme

4️⃣ Gelişmiş Özellikler (Sonraki Sürümler)

Long-to-wide ve wide-to-long reshape fonksiyonları

Renkli highlight ile QA sorunlarının tablo içinde gösterimi

Dry-run / preview modu

Config file ile preset temizleme/özetleme işlemleri

Opsiyonel terminal-only file picker (--terminal-picker flag ile)

5️⃣ Kullanıcı Akışı (UX)
Terminal Açılıyor → Menü:
[1] Load CSV/Excel (GUI file picker açılır)
[2] View Table
[3] Clean Data
[4] Summarize / Pivot
[5] Run QA Checks
[6] Export
[7] Exit


Dosya yükleme → GUI picker ile seçim

Tabloda navigasyon → ok tuşları

Temizleme / QA / Özetleme → space veya enter ile uygulama

Export → temizlenmiş veya özetlenmiş veri kaydetme

6️⃣ Tech Stack / Kütüphaneler
Amaç	Kütüphane / Tech	Not
CLI / TUI	github.com/charmbracelet/bubbletea	Interaktif terminal UI
GUI file picker	OS-native (osascript, powershell, zenity)	Platform bağımlı ama kullanıcı dostu
CSV	encoding/csv	Standart library
Excel	github.com/xuri/excelize/v2	XLSX read/write
Logging	github.com/sirupsen/logrus	Temiz log ve debug
String / Regex	strings / regexp	Column normalization, karakter temizleme
7️⃣ Proje Dizini (Boilerplate)
go-data-cleaner/
│
├── cmd/
│   ├── root.go            # Ana komut
│   ├── clean.go           # Temizleme komutları
│   ├── summarize.go       # Özetleme / pivot komutları
│   └── reshape.go         # Long-to-wide / wide-to-long
│
├── internal/
│   ├── cleaner/
│   │   └── cleaner.go
│   ├── summarizer/
│   │   └── summarizer.go
│   ├── reshaper/
│   │   └── reshaper.go
│   └── utils/
│       └── file.go        # CSV/XLSX read/write, GUI file picker
│
├── examples/              # Örnek CSV/XLSX dosyaları
├── go.mod
├── go.sum
└── README.md

8️⃣ MVP Başlatma Planı

TUI açılır, ana menü gösterilir

“Load CSV/Excel” seçildiğinde GUI file picker açılır

Kullanıcı dosyayı seçer → tablo Bubble Tea TUI’de gösterilir

Ok tuşları ile satır/sütun seçimi yapılır

Temizleme ve QA işlemleri uygulanır

Özet tablo oluşturulabilir

Export ile dosya kaydedilir

9️⃣ Performans ve Golang Avantajları

Büyük CSV/Excel dosyaları hızlı işlenebilir

Concurrency ile duplicate/missing check ve normalization parallel yapılabilir

Tek Go binary → kolay kurulum ve dağıtım

Bubble Tea TUI → modern, interaktif, profesyonel terminal arayüzü

🔟 MVP Başarı Ölçütleri

CSV/XLSX dosyası GUI file picker ile seçilip terminalde görüntülenebiliyor

Boş satır/sütun temizleme ve sütun normalize işlemleri uygulanabiliyor

Duplicate / missing value kontrolü yapılabiliyor

Temizlenmiş / özetlenmiş veri export edilebiliyor

Kullanıcı tüm işlemleri ok tuşları ve kısa navigasyon ile yapabiliyor