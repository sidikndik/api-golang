package model

type Product struct {
	Base
    Name        string  // Nama produk
    Description string  // Deskripsi produk
    Price       float64 // Harga produk
    Stock       int     // Stok produk
}
