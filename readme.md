## BASIC CRUD Travel App (Golang + Next.js)

Proyek ini adalah aplikasi latihan (bukan untuk produksi) yang dibuat untuk memperdalam pemahaman konsep dasar CRUD (Create, Read, Update, Delete) menggunakan backend Golang (Gin + GORM + PostgreSQL/compatible) dan frontend Next.js (App Router). Fokus utama: membangun alur end-to-end sederhana termasuk upload gambar ke Cloudinary.

> Catatan: Kode dan konfigurasi di sini bersifat edukatif, minimalis, dan dapat dikembangkan lebih lanjut sesuai kebutuhan nyata (penanganan error lebih rinci, pagination, auth, logging terstruktur, dsb.).

---

## 1. Tujuan Pembelajaran
1. Memahami struktur dasar backend Golang dengan Gin.
2. Menggunakan GORM untuk operasi database dasar pada entitas `Travel`.
3. Mengelola upload file (gambar) dan integrasi dengan Cloudinary.
4. Menyusun frontend sederhana dengan Next.js (App Router) untuk interaksi CRUD.
5. Menerapkan pola pemisahan layer: controller, model, request (DTO), routes, config.

---

## 2. Teknologi Utama
Backend:
- Golang + Gin
- GORM
- Cloudinary SDK (upload & destroy image)
- UUID (otomatis melalui default DB `gen_random_uuid()`)

Frontend:
- Next.js (App Router)
- Fetch API / form submission
- Styling dasar (CSS global)

Database: PostgreSQL

---

## 3. Struktur Direktori (Ringkas)
```
backend/
	main.go
	routes/travel-router.go
	controller/travel_controller/travel-controller.go
	model/travel-model.go
	request/travel-request.go
	database/migrations/000001_travel.up.sql
frontend/
	src/app/page.jsx (daftar data)
	src/app/travel/create/page.jsx
	src/app/travel/update/[id]/page.jsx
	src/app/travel/detail/[id]/page.jsx
```

---

## 4. Entity / Model: Travel
Field | Tipe | Keterangan
----- | ---- | ----------
id | uuid (string) | Primary key, generate otomatis oleh DB
name | string | Nama paket travel (required)
description | string | Deskripsi singkat (required)
photo | string (URL) | URL hasil upload Cloudinary
price | int | Harga (required)
created_at | timestamp | Otomatis oleh GORM
updated_at | timestamp | Otomatis oleh GORM
deleted_at | nullable timestamp | Soft delete (GORM)

Request Body (form-data untuk Create/Update dengan upload file):
- name (string)
- description (string)
- price (int)
- photo (file, optional saat update)

---

## 5. Endpoint Backend
Base URL (contoh lokal): `http://localhost:8000`

Method | Endpoint | Deskripsi | Body | Catatan
------ | -------- | --------- | ---- | -------
POST | /add-travel | Create data travel | form-data (name, description, price, photo) | Upload ke Cloudinary
GET | / | Ambil semua travel (aktif) | - | Urut created_at ASC
GET | /:id | Ambil detail travel by ID | - | 404 jika tidak ada
PUT | /update-travel/:id | Update travel | form-data (name, description, price, [photo]) | Jika photo baru, yang lama dihapus di Cloudinary
DELETE | /:id | Soft delete travel | - | Gunakan id UUID

Contoh Response Sukses (Create):
```json
{
	"Message": "Berhasil Menambahkan Data Travel",
	"Data": {
		"id": "e6b6f8b2-...",
		"name": "Bali Hemat",
		"description": "Liburan 3 hari 2 malam",
		"photo": "https://res.cloudinary.com/.../travel-Bali-Hemat-uuid.jpg",
		"price": 2500000,
		"CreatedAt": "2025-09-19T10:00:00Z",
		"UpdatedAt": "2025-09-19T10:00:00Z"
	}
}
```

Contoh Error Validasi:
```json
{
	"Message": "Bad Request",
	"Error": "Key: 'TravelRequest.Name' Error:Field validation for 'Name' failed on the 'required' tag"
}
```

---

## 6. Alur Upload & Update Gambar
1. Client mengirim form-data termasuk field `photo`.
2. Server membuka file dan inisialisasi Cloudinary (secure URL).
3. Penamaan file kustom: `travel-{name}-{uuid}`.
4. Saat update: jika ada foto baru dan foto lama ada, server memanggil `Destroy` Cloudinary menggunakan public ID yang diekstrak dari URL lama.
5. URL secure baru disimpan ke kolom `photo`.

Utility: `helper.ExtractPublicID()` (ekstraksi public ID dari URL untuk proses destroy).

---

## 7. Menjalankan Proyek (Lokal)
### 7.1 Prasyarat
- Go >= 1.21
- Node.js >= 18
- Database PostgreSQL berjalan & variabel env backend ter-set (DB URL, Cloudinary credentials).

### 7.2 Backend
Masuk folder `backend` lalu jalankan (contoh pakai air jika sudah diinstall):
```
air
```
Atau tanpa hot reload:
```
go run main.go
```
Default port (misal): 8000 (pastikan sesuai implementasi `main.go`).

### 7.3 Frontend
Masuk folder `frontend`:
```
npm install
npm run dev
```
Default Next.js dev server: `http://localhost:3000`.

---

## 8. Frontend (Ringkasan Halaman)
Halaman | Path | Fungsi
------- | ---- | ------
Daftar | `/` atau `/` di app router | Menampilkan list travel
Buat | `/travel/create` | Form tambah data + upload gambar
Detail | `/travel/detail/[id]` | Lihat 1 data
Update | `/travel/update/[id]` | Edit data + ganti foto opsional

Komponen Reusable: `components/table.jsx` untuk menampilkan tabel data.

---


## 9. Lisensi / Status
Tidak ada lisensi formal. Gunakan bebas untuk belajar. Jangan gunakan langsung untuk aplikasi produksi tanpa peningkatan signifikan.

---

## 10. Kesimpulan
**Secara keseluruhan, proyek ini menghadirkan kerangka kerja konseptual minimal yang menggambarkan siklus hidup data dalam arsitektur aplikasi web modern berbasis layanan terintegrasi. Penerapan pemisahan tanggung jawab antarlapisan—mulai dari antarmuka pengguna (Next.js), logika aplikasi dan orkestrasi I/O (Gin), hingga lapisan penyimpanan data (GORM + basis data relasional)—menyoroti bagaimana prinsip modularitas dan keterujian dapat diwujudkan meskipun dalam contoh sederhana.**

**Integrasi layanan penyimpanan eksternal (Cloudinary) menambah dimensi praktis dalam pengelolaan aset non-struktural, sekaligus menegaskan pentingnya pengelolaan siklus penuh sumber daya (unggah, referensi, hingga penghapusan). Meskipun implementasi ini masih terbatas—misalnya belum mencakup mekanisme autentikasi, observabilitas, pola ketahanan, atau pengamanan tingkat lanjut—kerangka kerja yang dihasilkan tetap menyediakan fondasi konseptual kuat untuk dikembangkan secara bertahap menjadi desain yang lebih matang dan tangguh.**

**Dengan demikian, kode proyek ini bukan hanya sekadar latihan sintaksis semata, melainkan juga artefak pedagogis yang menggambarkan keterkaitan antara desain skema data, kontrak antarmuka sederhana, dan pengelolaan status di lintas batas jaringan. Implementasi ini menjadi landasan penting untuk menginternalisasi paradigma rekayasa perangkat lunak terapan sebelum menghadapi kompleksitas sistem terdistribusi berskala besar.**

---