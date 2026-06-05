package model

import "gorm.io/gorm"

type Pengembalian struct {
    gorm.Model
    PeminjamanID     uint               `json:"peminjaman_id"`
    PetugasID        uint               `json:"petugas_id"`
    TglKembali       string             `json:"tgl_kembali"`

    // Relasi
    Peminjaman       Peminjaman         `json:"peminjaman,omitempty" gorm:"foreignKey:PeminjamanID"`
    Petugas          Petugas            `json:"petugas,omitempty" gorm:"foreignKey:PetugasID"`
    DetailPengembalian []DetailPengembalian `json:"detail_pengembalian,omitempty" gorm:"foreignKey:PengembalianID"`
}

type DetailPengembalian struct {
    gorm.Model
    PengembalianID    uint             `json:"pengembalian_id"`
    DetailPeminjamanID uint            `json:"detail_peminjaman_id"`
    KondisiBuku       int             `json:"kondisi_buku"`

    // Relasi
    DetailPeminjaman  DetailPeminjaman `json:"detail_peminjaman,omitempty" gorm:"foreignKey:DetailPeminjamanID"`
}