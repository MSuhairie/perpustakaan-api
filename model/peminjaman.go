package model

import "gorm.io/gorm"

type Peminjaman struct {
    gorm.Model
    AnggotaID      uint               `json:"anggota_id"`
    PetugasID      uint               `json:"petugas_id"`
    TglPinjam      string             `json:"tgl_pinjam"`
    TglJatuhTempo  string             `json:"tgl_jatuh_tempo"`
    Status         string             `json:"status"` // dipinjam, dikembalikan, terlambat

    // Relasi
    Anggota        Anggota            `json:"anggota,omitempty" gorm:"foreignKey:AnggotaID"`
    Petugas        Petugas            `json:"petugas,omitempty" gorm:"foreignKey:PetugasID"`
    DetailPeminjaman []DetailPeminjaman `json:"detail_peminjaman,omitempty" gorm:"foreignKey:PeminjamanID"`
    Denda          *Denda             `json:"denda,omitempty" gorm:"foreignKey:PeminjamanID"`
}

type DetailPeminjaman struct {
    gorm.Model
    PeminjamanID uint  `json:"peminjaman_id"`
    BukuID       uint  `json:"buku_id"`
    Jumlah       int   `json:"jumlah"`

    // Relasi
    Buku         Buku  `json:"buku,omitempty" gorm:"foreignKey:BukuID"`
}