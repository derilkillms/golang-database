# golang-database

- sql.DB di golang sebenarnya bukanlan sebuah koneksi ke database
- melainkan sebuah pool ke databasem atau dikenal dengan konsep Database Pooling
- Di dalam sql.DB, golang melakukan management koneksi ke database secara otomatis, Hal ini menjadikan kita tidak perlu melakukan management koneksi database secara manual
- Dengan kemampuan database pooling ini, kita bisa menentukan jumlah miniman dan maksimal koneksi yang dibuat oleh golang, sehingga tidak membanjiri koneksi ke database, karena biasanya ada batas maksimal koneksi yang bisa ditangani oleh database yang kita gunakan

|Method                           |	Keterangan                        |
| ------------------------------- | --------------------------------- |
|(DB) SetMaxIdleConns(number)     |	Pengaturan berapa jumlah koneksi minimal yang dibuat       |
|(DB) SetMaxOpenConns(number)     |	Pengaturan berapa jumlah koneksi maksimal yang dibuat      |
|(DB) SetConnMaxIdleTime(duration)|	Pengaturan berapa lama koneksi yang sudah tidak digunakan akan dihapus|
|(DB) SetConnMaxLifetime(duration)|	Pengaturan berapa lama koneksi boleh digunakan|

- Dalam buku Domain-Driven, Eric Evans menjelaskan bahwa “Repository is a mechanism for encapsulating storage, retrieval, and search behavior, which emulates a collection of objects”. 