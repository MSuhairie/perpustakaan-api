package repository

import (
    "perpustakaan-api/model"
    "gorm.io/gorm"
)

type PengembalianRepository interface {
    FindAll() ([]model.Pengembalian, error)
    FindByID(id string) (model.Pengembalian, error)
    Create(pengembalian model.Pengembalian) (model.Pengembalian, error)
}

type pengembalianRepository struct {
    db *gorm.DB
}

func NewPengembalianRepository(db *gorm.DB) PengembalianRepository {
    return &pengembalianRepository{db}
}

func (r *pengembalianRepository) FindAll() ([]model.Pengembalian, error) {
    var list []model.Pengembalian
    err := r.db.
        Preload("Peminjaman").
        Preload("Peminjaman.Anggota").
        Preload("Peminjaman.DetailPeminjaman").
        Preload("Peminjaman.DetailPeminjaman.Buku").
        Preload("Petugas").
        Preload("DetailPengembalian").
        Preload("DetailPengembalian.DetailPeminjaman").
        Preload("DetailPengembalian.DetailPeminjaman.Buku").
        Find(&list).Error
    return list, err
}

func (r *pengembalianRepository) FindByID(id string) (model.Pengembalian, error) {
    var pengembalian model.Pengembalian
    err := r.db.
        Preload("Peminjaman").
        Preload("Peminjaman.Anggota").
        Preload("Peminjaman.DetailPeminjaman").
        Preload("Peminjaman.DetailPeminjaman.Buku").
        Preload("Petugas").
        Preload("DetailPengembalian").
        Preload("DetailPengembalian.DetailPeminjaman").
        Preload("DetailPengembalian.DetailPeminjaman.Buku").
        First(&pengembalian, id).Error
    return pengembalian, err
}

func (r *pengembalianRepository) Create(pengembalian model.Pengembalian) (model.Pengembalian, error) {
    err := r.db.Create(&pengembalian).Error
    return pengembalian, err
}