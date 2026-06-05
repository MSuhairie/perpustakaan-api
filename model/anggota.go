package model

import "gorm.io/gorm"

type Anggota struct {
    gorm.Model
    Nama       string      `json:"nama"`
    Alamat     string      `json:"alamat"`
    NoHP       string      `json:"no_hp"`
    TglDaftar  string      `json:"tgl_daftar"`
    Status     string      `json:"status"`
    Peminjaman []Peminjaman `json:"peminjaman,omitempty" gorm:"foreignKey:AnggotaID"`
}