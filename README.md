# Jagona Gym Project

---
Projek ini adalah projek untuk memenuhi tugas mata kuliah Projek RPL Kelas AI

## Memulai

Sebelum memulai pastikan untuk **mengunduh [Go](https://go.dev/doc/install)** versi **1.19** atau yang lebih baru.

1. Klon proyek ini ke lokal
    ```
    $ git clone https://github.com/akramfirmansyah/jagona-gym.git
    $ cd jagona-gym
    ```

2. Unduh semua package yang dibutuhkan
    ```
    $ go get
    ```

3. Buat file `.env` dan salin isi `.env.example` ke file `.env` yang telah dibuat.

4. Buatlah database baru. (Dalam proyek ini menggunakan **[MySQL](https://www.mysql.com/)** sebagai DBMS)

5. Sesuaikan konfigurasi Database di file `.env` (terutama `DB_HOST`, `DB_PASS` dan `DB_NAME`) dengan database yang telah dibuat

6. Jalankan program
    ```
    $ go run main.go
    ```
7. Buka [localhost](http://127.0.0.1:8080) dan akan muncul pesan `Jagona Gym Project`

8. **(Jika Terjadi Error)** silahkan berkabar di Grup.