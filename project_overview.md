# OgiTech Handover: SIGAP PANGAN (SULTAN) Project Overview

Selamat datang tim **OgiTech**! Dokumen ini disusun khusus untuk memberikan Anda pemahaman komprehensif mengenai *repository* SIGAP SULTAN (Sistem Informasi Harga Pangan Sulawesi Selatan - Bank Indonesia) peninggalan vendor sebelumnya (Sentech).

Analisis ini dibagi menjadi tiga perspektif utama:

---

## 1. Perspektif *Engineering* (Arsitektur & Teknologi)

Secara arsitektur, proyek ini tidak dibangun sebagai sistem *Monolith*, melainkan menggunakan pendekatan terdistribusi (*Decoupled / Microservices-oriented*). Hal ini dilakukan untuk memisahkan beban kerja antara portal publik dan *back-office*.

### A. SPBI-BACKEND (Core Engine)
- **Tech Stack**: Golang (Go v1.20) + Fiber Framework v2.
- **Database**: PostgreSQL (Via `pgx` dan `sqlx`) + Redis (untuk Caching).
- **Peran**: Menjadi "Otak" utama dari sistem. Backend ini murni berupa RESTful API yang melayani seluruh kalkulasi berat seperti perbandingan harga antar wilayah, penentuan status defisit/aman berdasarkan stok, hingga *Month-to-Month (MTM) calculations*.
- **Otentikasi**: Menggunakan JWT.

### B. SPBI-CMS (Back-Office / Admin Panel)
- **Tech Stack**: PHP (>= 8.1) + Laravel 10 + Bootstrap / SB Admin 2.
- **Database**: MySQL.
- **Peran**: Aplikasi *Server-Side Rendering* mandiri yang hanya diakses oleh pegawai/staf internal. Menyimpan data kredensial pegawai (MySQL) terpisah dari data pangan (PostgreSQL). Bertugas sebagai pintu masuk data (*Data Ingestion*), di mana admin mengunggah file Excel berisi harga/stok harian, lalu CMS ini akan meneruskannya (via HTTP POST) ke SPBI-BACKEND.

### C. SPBI-FRONTEND (Public / Executive Dashboard)
- **Tech Stack**: Angular 17 + TypeScript + Tailwind CSS v3.
- **Peran**: *Single Page Application* (SPA) yang sangat responsif. Aplikasi ini hanya *memakan* (consume) API dari Golang. Dibangun dengan dukungan PWA (*Progressive Web App*) sehingga bisa diinstal layaknya aplikasi *native* di perangkat pengguna.

> [!TIP]
> **Dockerization**: Ketiga sistem ini, beserta ketiga databasenya, baru saja saya konfigurasi ulang menjadi 1 environment menggunakan `docker-compose.yml`. Anda bisa langsung mem-build dan me-running semuanya secara lokal.

---

## 2. Perspektif *Business Flow* (Alur Bisnis)

Tujuan utama dari sistem SIGAP SULTAN ini adalah untuk membantu **Bank Indonesia (BI) dan Pemerintah Provinsi Sulawesi Selatan** dalam mengambil kebijakan pengendalian inflasi dan ketahanan pangan.

### Masalah yang Diselesaikan:
Ketimpangan harga dan stok komoditas (seperti Beras, dll) di berbagai Kabupaten/Kota (Makassar, Palopo, Wajo, Bone, Parepare).

### Indikator Utama yang Dipantau:
1. **Neraca Ketersediaan (Stok vs Kebutuhan)**:
   Sistem mengkategorikan wilayah ke dalam 4 *Tier* (*Threshold*):
   - **Aman** (Warna Hijau - `#05603A`)
   - **Waspada** (Warna Kuning - `#E4B701`)
   - **Rentan** (Warna Oranye - `#FF6711`)
   - **Defisit** (Warna Merah - `#B11016`): Terjadi jika Kebutuhan lebih besar dari Ketersediaan.
2. **Perbandingan Harga (Price Tracker)**:
   Melacak harga MTM (*Month to Month*), serta membandingkan secara *real-time* harga sebuah komoditas di tingkat Kota vs Provinsi vs Nasional.

---

## 3. Perspektif *User Flow* (Alur Pengguna)

Sistem ini melayani dua jenis pengguna dengan *journey* yang sama sekali berbeda:

### A. Alur Admin / Staf (Menggunakan CMS Laravel)
1. **Login**: Admin masuk ke URL CMS (port `8000`).
2. **Manajemen Akun**: Super-Admin dapat membuat akun baru dan memberikan "Jabatan" spesifik.
3. **Data Ingestion (Upload)**: Tugas utama staf harian adalah masuk ke menu **"Unggah Neraca"** atau **"Unggah Harga"**.
4. **Validasi**: Admin memilih file Excel (.xlsx / .csv). Sistem CMS memvalidasi isi baris tersebut.
5. **Sinkronisasi**: Jika valid, CMS mem-parsing data tersebut dan melakukan proses sinkronisasi API ke server Backend Golang.
6. **Log**: Admin melihat menu "Unggah Log" untuk memastikan tidak ada data yang gagal ter-input di database sentral.

### B. Alur Publik / Pimpinan Eksekutif (Menggunakan Frontend Angular)
1. **Akses Dashboard**: Pimpinan BI/Pemprov membuka URL Frontend (port `4200`).
2. **Visualisasi Peta**: Mereka langsung disambut oleh Peta Sulawesi Selatan yang diwarnai secara dinamis (Misal: Kabupaten Bone menyala merah karena stok beras bulan ini sedang *Defisit*).
3. **Drill-down Data**: Pimpinan mengklik Kabupaten Bone. Muncul *sidebar/modal* interaktif berisi grafik historikal harga beras dan tren pasokan dalam beberapa bulan terakhir.
4. **Perbandingan (Komparasi)**: Pimpinan dapat menggunakan fitur filter untuk membandingkan langsung harga di Kota Makassar vs Kabupaten Wajo guna menentukan ke mana subsidi/intervensi pasar harus difokuskan.

---

## 4. Perspektif Fitur & Database (Rincian Lengkap)

### A. Rincian Aktor & Hak Akses Fitur
Sistem dibagi menjadi dua aplikasi, sehingga fiturnya pun terbagi:

#### 1. SPBI-FRONTEND (Akses: Publik / Pimpinan)
Semua fitur di sini bersifat **Read-Only** (Hanya melihat data):
- **Peta Interaktif Sulawesi Selatan**: Menampilkan map dengan indikator warna (hijau/kuning/merah) berdasarkan rasio stok vs kebutuhan.
- **Grafik Komparasi Harga MTM**: Membandingkan inflasi/deflasi harga bulan ini vs bulan lalu.
- **Grafik Agregasi Wilayah**: Membandingkan harga rata-rata komoditas di tingkat Kota vs Provinsi vs Nasional.
- **Filter Dinamis**: Memfilter data berdasarkan rentang tanggal, jenis komoditas, atau kota spesifik.

#### 2. SPBI-CMS (Akses: Admin / Staf Terautentikasi)
Fitur di sini bersifat **Manajemen / Operasional**:
- **Modul Autentikasi**: `Login`, `Logout`, `Lupa Password (Forgot)`, `Reset Password`.
- **Modul Manajemen Jabatan (Roles)**: `Lihat Daftar Jabatan`, `Tambah Jabatan Baru`, `Lihat Detail`, `Hapus Jabatan`.
- **Modul Manajemen Pengguna (Users)**: `Lihat Daftar Staf`, `Tambah Akun Baru`, `Edit Data Staf`, `Hapus Akun`.
- **Modul Manajemen Data Pangan (Data Ingestion)**:
  - `Unggah Neraca`: Form untuk mengunggah file Excel berisi laporan Stok & Kebutuhan bulanan.
  - `Unggah Harga`: Form untuk mengunggah file Excel berisi laporan fluktuasi Harga Pasar bulanan/harian.
  - `Riwayat Unggahan (Upload Log)`: Melihat tabel log kesuksesan sinkronisasi file Excel ke sistem backend utama.

### B. Struktur Penyimpanan Database
Sistem ini menggunakan penyimpanan data secara *hybrid* di dua tempat.

#### 1. Disimpan di MySQL (Milik CMS Laravel)
Fokus pada data pengguna operasional *Back-Office*:
- **`users`**: Menyimpan email, nama, password (terenkripsi *hash bcrypt*), dan relasi `jabatan_id`.
- **`password_reset_tokens`**: Menyimpan token unik sementara jika ada admin yang lupa password.

#### 2. Disimpan di PostgreSQL (Milik API Backend Golang)
Fokus pada big data / *Data Warehouse* utama:
- **Tabel Referensi (Master Data)**:
  - `tm_province` & `tm_city`: Menyimpan data geografis Provinsi dan Kota/Kabupaten.
  - `tm_commodity`: Menyimpan jenis komoditas (Beras, Minyak, Gula, dll).
  - `tm_position`: Menyimpan sinkronisasi referensi Jabatan/Roles dari CMS.
  - `tm_jenis_informasi`: Menyimpan referensi jenis tipe data pelaporan.
- **Tabel Transaksional (Inti Sistem)**:
  - `price`: Menyimpan angka harga suatu komoditas di kota tertentu pada tanggal tertentu.
  - `neraca`: Menyimpan *record* ketersediaan stok vs tingkat konsumsi harian.
  - `tx_file_upload_history`: Menyimpan metadata log unggahan file Excel (Nama file, tanggal, nama pengunggah, status).
- **Tabel Media**:
  - `assets`: Menyimpan referensi URL gambar eksternal pendukung (Logo kota, icon sembako, dll).

Semua perhitungan *threshold* (*"Apakah kota ini Defisit atau Aman?"*) tidak disimpan mentah-mentah di database, melainkan **dihitung secara cerdas (*on-the-fly*)** oleh kalkulator *Backend Golang* ketika di-*request*.

---

## 5. Rencana Tindak Lanjut: Apa yang Akan OgiTech Lakukan?

Fokus utama OgiTech adalah menyelesaikan *The Last Mile* dari proyek ini, yaitu **Membangun Frontend Executive Dashboard (Peta Interaktif)** yang belum diselesaikan oleh vendor sebelumnya. 

Adapun rincian tugas (*Developer Task Breakdown*) yang akan dikerjakan tim *Frontend Engineer* OgiTech meliputi:

1. **Perancangan UI/UX Premium (*UI Engineering*)**: 
   - **Tugas:** Merancang antarmuka dashboard bergaya *dark-mode* dan *glassmorphism* menggunakan *Tailwind CSS v3*.
   - **Output:** Tampilan responsif yang terlihat profesional di layar besar (presentasi eksekutif) maupun *mobile*.

2. **Pengembangan Peta Geospasial (GIS) (*Component: MapView*)**: 
   - **Tugas:** Mengintegrasikan library pemetaan *Open-Source* (seperti `Leaflet.js` via `@asymmetrik/ngx-leaflet` atau `Mapbox GL JS`) ke dalam kerangka *Angular 17*.
   - **Output:** *Render polygon* 24 Kabupaten/Kota di Sulawesi Selatan (butuh file `.geojson` batas wilayah administratif Sulsel).

3. **Pewarnaan Dinamis / Choropleth (*State Management*)**: 
   - **Tugas:** Menulis *service* di Angular (RxJS) untuk memetakan data numerik ke warna polygon di peta.
   - **Output:** Peta otomatis berubah warna (Aman = Hijau, Defisit = Merah) berdasarkan logika matematika dari rasio "Ketersediaan vs Kebutuhan". Dilengkapi fitur *Hover/Tooltip* untuk melihat angka pasti.

4. **Integrasi Grafik Data / Charts (*Component: AnalyticsSidebar*)**: 
   - **Tugas:** Membangun *Reusable Components* untuk visualisasi data menggunakan library chart (seperti `Chart.js` via `ng2-charts` atau `ApexCharts`).
   - **Output:** Menampilkan diagram batang untuk stok *real-time* dan grafik garis untuk melacak tren inflasi harga (Month-to-Month) secara dinamis saat suatu kota di-klik.

5. **Konsumsi REST API (*Data Fetching*)**: 
   - **Tugas:** Membuat `HttpClient` services di Angular untuk menembak *endpoint* Golang yang sudah jadi (contoh: `GET /api/v1/neraca/map-summary`).
   - **Output:** Aliran data JSON dari *backend* berhasil diparsing dan dirender tanpa *delay* ke komponen visual (Peta dan Grafik).
   - ** CATATAN PENTING:** *Tidak perlu lagi merancang database, ERD, atau membuat API dari awal. Vendor sebelumnya (Sentech) sudah menyelesaikan seluruh "mesin" Golang dan CMS untuk input data. Tugas utama kita (OgiTech) murni mem-parsing data dari endpoint API yang sudah tersedia tersebut dan menyajikannya menjadi visualisasi interaktif yang elegan.*

---