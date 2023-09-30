package model

type User struct {
    Base
	Name string
    Username string // Nama pengguna
    Email    string // Alamat email pengguna
    Password string // Kata sandi pengguna
}


func (User) TableName() string {
    return "user"
}