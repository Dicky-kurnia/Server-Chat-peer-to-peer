Chat App
Aplikasi chat peer-to-peer yang menggunakan websocket dan dapat menyimpan dan mengambil pesan secara realtime.

Fitur
User dapat mendaftar dan masuk dengan email dan password
Server dapat menerima pesan dari websocket dan menyimpannya ke database
Chat hanya dapat dilakukan antar user, bukan group chat
User dapat mencari riwayat chat dari database
REST API yang aman untuk Chat
REST API untuk mengambil chat peer-to-peer
Menggunakan kueri SQL mentah, bukan ORM
Skrip migrasi untuk membuat tabel di database
User dapat mengirim pesan meskipun penerima offline dan jika penerima online kembali, pesan akan diterima oleh penerima
Menggunakan RabbitMQ untuk menangani antrian untuk pengguna offline
Unit test

Teknologi yang digunakan
Backend: Go
API Framework: Go Fiber
PostgreSQL interfaces for Go: database/sql
Database: PostgreSQL
Docker

Instalasi
Clone repository ini
git clone https://github.com/yourusername/chat-app.git
Pindah ke direktori chat-app

cd chat-app
Jalankan perintah docker-compose untuk membuat dan menjalankan container
docker-compose up --build
Buka browser dan akses http://localhost:3000 untuk mengakses aplikasi.
Cara Penggunaan
Register dengan email dan password
Login dengan email dan password yang sudah didaftarkan
Kirim pesan ke user lain
Lihat riwayat chat

Unit Test
Untuk menjalankan unit test, jalankan perintah berikut di dalam direktori aplikasi:
go test ./...

Catatan
Pastikan semua dependensi sudah terinstall sebelum menjalankan aplikasi
Pastikan konfigurasi database dan rabbitmq sudah sesuai dengan yang digunakan pada aplikasi
