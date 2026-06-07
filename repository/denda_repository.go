package repository

import (
    "perpustakaan-api/model"
    "gorm.io/gorm"
)

type DendaRepository interface {
    FindByPeminjaman(peminjamanID uint) (model.Denda, error)
    Create(denda model.Denda) (model.Denda, error)
}

type dendaRepository struct {
    db *gorm.DB
}

func NewDendaRepository(db *gorm.DB) DendaRepository {
    return &dendaRepository{db}
}

func (r *dendaRepository) FindByPeminjaman(peminjamanID uint) (model.Denda, error) {
    var denda model.Denda
    err := r.db.Where("peminjaman_id = ?", peminjamanID).First(&denda).Error
    return denda, err
}

func (r *dendaRepository) Create(denda model.Denda) (model.Denda, error) {
    err := r.db.Create(&denda).Error
    return denda, err
}