# Online Store RESTful API

Repositori ini berisi implementasi sederhana dari online store RESTful API yang dikembangkan dengan bahasa Golang menggunakan framework Fiber dan database MySQL. API ini menyediakan endpoint untuk mengelola produk, keranjang belanja, proses chekout dan juga register dan juga login.

## Syarat-Syarat

* Install Golang 
* MySQL
* Visual Studio Code
* Postman

## Clone Repositori Dibawah Ini
```bash
git clone https://github.com/sigit-khoirun-nizam/go-fiber-onlinestore.git
```

## Buka Direktori Project
```bash
cd go-fiber-onlinestore
```

## Konfigurasi database MySQL
* Buat database baru dengan nama `synapsis`.
* Perbarui detail koneksi database difungsi `InitDatabase` dalam folder utils lalu filenya bernama `database.go`.

## Jalankan Aplikasi
```bash
go run main.go
```

Server API akan berjalan di `http://localhost:3000`.

# Endpoint API

* Mendapatkan produk berdasarkan kategori
```bash
GET /products/{category}
```

* Tambah produk 
```bash
POST /products
```

> Body Request
```json
{
	"name": "Nama Produk",
	"category": "Kategori Produk",
	"price": 20000
}
```

* Mendapatkan semua produk
```bash
GET /products
```

* Update produk 
```bash
PUT /products/{id}
```

> Body Request
```json
{
	"name": "Nama Produk Baru",
	"category": "Kategori Produk Baru",
	"price": 20000
}
```

* Hapus produk 
```bash
DELETE /products/{id}
```

* Tambah produk ke keranjang
```bash
POST /cart/add
```

> Body Request
```json
{
    "product_id"    : 1,
    "quantity"      : 2
}
```

* Mendapatkan produk ke keranjang
```bash
GET /cart
```

* Hapus produk dari keranjang
```bash
DELETE /cart/{id}
```

* Checkout dan pembayaran
```bash
POST /checkout
```

* Register user
```bash
POST /register
```

> Body Request
```json
{
    "username": "Nizam",
    "password": "rahasia123"
}
```

* Login user
```bash
POST /login
```

> Body Request
```json
{
    "username": "Nizam",
    "password": "rahasia123"
}

```

