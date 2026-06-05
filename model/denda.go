package model

import "gorm.io/gorm"

type Denda struct {
    gorm.Model
    PeminjamanID uint   `json:"peminjaman_id"`
    JumlahDenda  int    `json:"jumlah_denda"`
    Keterangan   string `json:"keterangan"`
}