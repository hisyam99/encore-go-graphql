# Panduan untuk Menjalankan Fix Blog Status

## Masalah yang Diperbaiki
- Status blog disimpan dalam uppercase (contoh: "PUBLISHED") sehingga hook `BeforeCreate` dan `BeforeUpdate` tidak dapat mengisi field `PublishedAt`
- Query `publishedBlogs` tidak mengembalikan hasil karena filter mencari status lowercase ("published") tapi data tersimpan uppercase

## Perubahan yang Dilakukan

### 1. Menambahkan fungsi normalizeBlogStatus
- Fungsi ini menormalkan status blog ke lowercase
- Memastikan konsistensi antara "PUBLISHED" -> "published" dan "DRAFT" -> "draft"

### 2. Update di BlogService
- `CreateBlog`: Status dinormalisasi sebelum disimpan
- `UpdateBlog`: Status dinormalisasi saat update
- `ListBlogsByStatus`: Status dinormalisasi saat query

### 3. Menambahkan fungsi FixBlogStatus
- Di repository: `FixBlogStatus()` untuk mengupdate data yang sudah ada
- Di service: wrapper untuk memanggil repository method

## Cara Menjalankan Fix

### Opsi 1: Manual SQL (Paling Mudah)
Jalankan query SQL berikut di database Anda:

```sql
-- Update status to lowercase
UPDATE blogs SET status = LOWER(status) WHERE status != LOWER(status);

-- Set published_at for published blogs that don't have it set
UPDATE blogs SET published_at = COALESCE(published_at, updated_at) WHERE status = 'published' AND published_at IS NULL;
```

### Opsi 2: Menggunakan Script Go
1. Edit file `scripts/fix_blog_status.go` dan sesuaikan connection string database
2. Jalankan: `go run scripts/fix_blog_status.go`

### Opsi 3: Melalui GraphQL/API
Anda bisa menambahkan mutation baru untuk menjalankan fix ini melalui API.

## Verifikasi
Setelah menjalankan fix:
1. Cek bahwa semua status blog sekarang lowercase: `SELECT DISTINCT status FROM blogs;`
2. Cek bahwa blog published memiliki published_at: `SELECT id, title, status, published_at FROM blogs WHERE status = 'published';`
3. Test query `publishedBlogs` seharusnya mengembalikan hasil

## Hasil yang Diharapkan
- Semua blog dengan status "PUBLISHED" akan menjadi "published"
- Semua blog dengan status "DRAFT" akan menjadi "draft"
- Blog published yang sebelumnya published_at-nya null sekarang akan terisi
- Query `publishedBlogs` akan mengembalikan hasil yang benar