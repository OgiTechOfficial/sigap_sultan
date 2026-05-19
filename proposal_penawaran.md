# PROPOSAL PENGEMBANGAN FITUR GEOGRAPHIC INFORMATION SYSTEM (GIS) & EXECUTIVE DASHBOARD
**Sistem Informasi Harga Pangan (SIGAP SULTAN) - Bank Indonesia**
**Oleh: OgiTech**

---

## 1. Latar Belakang & Analisis Kondisi Saat Ini
Berdasarkan audit kode sumber (source code) terkini pada repositori `SPBI-FRONTEND`, kami menemukan bahwa **fitur Peta Interaktif dan Dashboard Publik sama sekali belum dikembangkan**. File `dashboard.component.html` saat ini masih berupa *template* kosong bawaan Angular (`<p>dashboard works</p>`).

Vendor sebelumnya baru menyelesaikan:
1. **CMS Admin (Laravel)**: Untuk manajemen user dan unggah data.
2. **Backend API (Golang)**: Endpoint data spasial dan algoritma kalkulasi defisit/aman.

Oleh karena itu, OgiTech hadir untuk **menyelesaikan jembatan terakhir (The Last Mile)** dari proyek ini, yaitu membangun *Executive Dashboard* berbasis *Geographic Information System (GIS)* yang memukau, responsif, dan interaktif menggunakan Angular.

---

## 2. Ruang Lingkup Pekerjaan (Scope of Work)

Meskipun secara konseptual ini adalah "1 fitur utama", pengembangan Peta Interaktif (GIS) membutuhkan kompleksitas *Frontend Engineering* tingkat tinggi. Pekerjaan OgiTech mencakup:

1. **Pengembangan Peta GIS Interaktif Sulawesi Selatan**:
   - Pemetaan polygon (*GeoJSON*) untuk 24 Kabupaten/Kota di Sulawesi Selatan.
   - Pewarnaan dinamis (*Choropleth Map*) berdasarkan *Threshold API*:
     - 🟢 **Aman**: Ketersediaan > Kebutuhan
     - 🟡 **Waspada / 🟠 Rentan**: Ketersediaan mendekati batas konsumsi
     - 🔴 **Defisit**: Kebutuhan melampaui ketersediaan
   - *Hover & Tooltip*: Menampilkan pop-up detail angka stok dan kebutuhan saat kursor diarahkan ke suatu kota.
2. **Pengembangan Panel Grafik Analitik (Charts)**:
   - Integrasi grafik *Month-to-Month (MTM)* untuk inflasi komoditas.
   - Integrasi perbandingan harga level Kota vs Provinsi vs Nasional.
3. **Integrasi REST API Golang**:
   - Menghubungkan *Frontend* Angular dengan Endpoint Golang yang sudah ada (misal: `NERACA_KETERSEDIAAN_BY_COMMODITY_MAP`).
4. **UI/UX Enhancement**:
   - Penerapan desain *Glassmorphism* dan *Dark/Light Mode* menggunakan Tailwind CSS agar terlihat premium di layar pimpinan (Executive View).

---

## 3. Contoh Tampilan (Mockup Prototipe)
Sebagai gambaran awal, berikut adalah standar desain premium yang akan dibangun oleh tim OgiTech untuk Pimpinan Eksekutif BI:

![Mockup Dashboard SIGAP SULTAN](C:\Users\Workplus\.gemini\antigravity\brain\e9eb9212-4556-4349-be9a-f48339e6d597\sigap_sultan_mockup_1779109791956.png)

---

## 4. Estimasi Waktu Pengerjaan (Timeline)
Total waktu pengembangan adalah **4 Minggu (1 Bulan)** dengan rincian:
- **Minggu 1**: UI/UX Design & Setup *Map Engine* (Leaflet.js / Mapbox).
- **Minggu 2**: Integrasi Polygon Peta Sulsel & Layouting Komponen Angular.
- **Minggu 3**: Integrasi API Golang, Interaktivitas Data, dan Pembuatan Grafik.
- **Minggu 4**: *Quality Assurance (QA)*, *Bug Fixing*, & Deployment.

---

## 5. Rencana Anggaran Biaya (RAB)
Mengingat kompleksitas integrasi peta geografis (GIS) dan ekspektasi tampilan *premium* standar Bank Indonesia, berikut adalah estimasi penawaran:

| Deskripsi Pekerjaan | Kuantitas / Peran | Estimasi Biaya |
| :--- | :--- | :--- |
| UI/UX Designer | 1 Orang (1 Bulan) | Rp 8.000.000 |
| Frontend Engineer (Angular & GIS Expert) | 2 Orang (1 Bulan) | Rp 25.000.000 |
| QA Tester & Deployment | 1 Orang (2 Minggu) | Rp 5.000.000 |
| **Sub-Total Pengembangan** | | **Rp 38.000.000** |
| Garansi Bug & Maintenance (3 Bulan) | | Rp 7.000.000 |
| **TOTAL BIAYA (Exclude PPN)** | | **Rp 45.000.000** |

*(Catatan: Harga di atas adalah estimasi yang bisa dinegosiasikan sesuai dengan pagu anggaran dari instansi pemberi kerja).*
