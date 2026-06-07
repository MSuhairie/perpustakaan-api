package repository

import (
    "perpustakaan-api/model"
    "gorm.io/gorm"
)

type PeminjamanRepository interface {
    FindAll() ([]model.Peminjaman, error)
    FindByID(id string) (model.Peminjaman, error)
    FindByAnggota(anggotaID string) ([]model.Peminjaman, error)
    Create(peminjaman model.Peminjaman) (model.Peminjaman, error)
    Update(peminjaman model.Peminjaman, input model.Peminjaman) (model.Peminjaman, error)
    Delete(peminjaman model.Peminjaman) error
}

type peminjamanRepository struct {
    db *gorm.DB
}

func NewPeminjamanRepository(db *gorm.DB) PeminjamanRepository {
    return &peminjamanRepository{db}
}

func (r *peminjamanRepository) FindAll() ([]model.Peminjaman, error) {
    var list []model.Peminjaman
    err := r.db.
        Preload("Anggota").
        Preload("Petugas").
        Preload("DetailPeminjaman").
        Preload("DetailPeminjaman.Buku").
        Preload("Denda").
        Find(&list).Error
    return list, err
}

func (r *peminjamanRepository) FindByID(id string) (model.Peminjaman, error) {
    var peminjaman model.Peminjaman
    err := r.db.
        Preload("Anggota").
        Preload("Petugas").
        Preload("DetailPeminjaman").
        Preload("DetailPeminjaman.Buku").
        Preload("Denda").
        First(&peminjaman, id).Error
    return peminjaman, err
}

func (r *peminjamanRepository) FindByAnggota(anggotaID string) ([]model.Peminjaman, error) {
    var list []model.Peminjaman
    err := r.db.
        Preload("Anggota").
        Preload("Petugas").
        Preload("DetailPeminjaman").
        Preload("DetailPeminjaman.Buku").
        Where("anggota_id = ?", anggotaID).
        Find(&list).Error
    return list, err
}

func (r *peminjamanRepository) Create(peminjaman model.Peminjaman) (model.Peminjaman, error) {
    err := r.db.Create(&peminjaman).Error
    return peminjaman, err
}

func (r *peminjamanRepository) Update(peminjaman model.Peminjaman, input model.Peminjaman) (model.Peminjaman, error) {
    err := r.db.Model(&peminjaman).Updates(input).Error
    return peminjaman, err
}

func (r *peminjamanRepository) Delete(peminjaman model.Peminjaman) error {
    return r.db.Delete(&peminjaman).Error
}