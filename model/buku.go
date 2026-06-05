package model

import "gorm.io/gorm"

type Buku struct {
	gorm.Model
	KategoriID  uint `json:"kategori_id"`
	RakID       uint `json:"rak_id"`
	Judul 		string `json:"judul"`
	Penulis 	string `json:"penulis"`
	Penerbit 	string `json:"penerbit"`
	TahunTerbit string  `json:"tahun_terbit"`
	ISBN 		string `json:"isbn"`
	Stok 		int `json:"stok"`
	Foto        string   `json:"foto"`
	Kategori 	Kategori `json:"kategori,omitempty" gorm:"foreignKey:KategoriID"`
	Rak      	Rak      `json:"rak,omitempty" gorm:"foreignKey:RakID"`
}