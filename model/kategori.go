package model

import "gorm.io/gorm"

type Kategori struct {
    gorm.Model
    NamaKategori string `json:"nama_kategori"`
    Buku         []Buku `json:"buku,omitempty" gorm:"foreignKey:KategoriID"`
}