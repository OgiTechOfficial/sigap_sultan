# Panduan Menjalankan SIGAP SULTAN via Docker

Semua file konfigurasi telah berhasil dibuat. Anda sekarang memiliki infrastruktur containerization lengkap untuk menjalankan frontend, backend, CMS, dan seluruh database secara serentak.

## File yang Ditambahkan & Dimodifikasi

1. **Modifikasi Kode Backend**:
   - `SPBI-BACKEND-main/src/config/db_config.go`: Menambahkan `os.Getenv("REDIS_HOST")` agar *backend* bisa terkoneksi dengan kontainer Redis dinamis, bukan *hardcode* ke `localhost`.

2. **File Dockerfile Baru**:
   - `SPBI-BACKEND-main/Dockerfile`: Konfigurasi *build* untuk Golang (API Backend).
   - `SPBI-CMS-main/Dockerfile`: Konfigurasi untuk PHP/Laravel CMS.
   - `SPBI-FRONTEND-main/Dockerfile`: Konfigurasi untuk Angular Frontend.

3. **File Orchestrator**:
   - `docker-compose.yml` di folder `sigap_sultan`: Menghubungkan seluruh 6 layanan (*backend, frontend, cms, postgres, mysql, redis*) menjadi satu kesatuan di dalam satu *network* (`sigap_network`).

---

## Langkah Mengeksekusi Proyek

> [!IMPORTANT]
> Pastikan aplikasi Docker Desktop (atau engine Docker) sudah berjalan di komputer Windows Anda sebelum mengeksekusi perintah di bawah ini.

### 1. Buka Terminal
Buka terminal (Command Prompt atau PowerShell) lalu arahkan ke direktori root proyek:
```bash
cd "d:\Projek BI\sigap-sultan-all-source-code\sigap_sultan"
```

### 2. Jalankan Perintah Build & Up
Jalankan satu perintah pamungkas ini. Ini akan memakan waktu lumayan lama saat pertama kali dijalankan karena harus men-*download* berbagai dependensi OS (Ubuntu/Alpine), Golang, Node.js, dan PHP.

```bash
docker-compose up --build -d
```
*(Catatan: flag `-d` digunakan agar proses berjalan di background)*.

### 3. Akses Aplikasi
Jika semua container sudah berstatus *Up* (Bisa dicek dengan perintah `docker-compose ps`), Anda dapat mengakses sistem di browser:

- **Frontend (Angular)**: [http://localhost:4200](http://localhost:4200)
- **CMS (Laravel)**: [http://localhost:8000](http://localhost:8000)
- **Backend API (Golang)**: [http://localhost:8080](http://localhost:8080)

> [!TIP]
> Jika Anda mengalami error terkait database *(Misalnya migrasi tabel kosong)* saat mengakses CMS/Backend, Anda dapat masuk ke dalam container masing-masing menggunakan Terminal Docker Desktop untuk menjalankan `php artisan migrate` atau proses *seeding* database.

### 4. Mematikan Server
Jika Anda ingin mematikan semua server, cukup jalankan perintah:
```bash
docker-compose down
```
