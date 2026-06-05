package repository

import (
	"perpustakaan-api/model"

	"gorm.io/gorm"
)

type AnggotaRepository interface {
	FindAll() ([]model.Anggota, error)
	FindByID(id string) (model.Anggota, error)
	FindByJudul(keyword string) ([]model.Anggota, error)
	Create(anggota model.Anggota) (model.Anggota, error)
	Update(anggota model.Anggota, input model.Anggota) (model.Anggota, error)
	Delete(anggota model.Anggota) error
}

type anggotaRepository struct {
	db *gorm.DB
}

func NewAnggotaRepository(db *gorm.DB) AnggotaRepository {
	return &anggotaRepository{db}
}

func (r *anggotaRepository) FindAll() ([]model.Anggota, error) {
	var anggotaList []model.Anggota
	err := r.db.Preload("Peminjaman").Find(&anggotaList).Error
	return anggotaList, err
}

func (r *anggotaRepository) FindByID(id string) (model.Anggota, error) {
    var anggota model.Anggota
    err := r.db.Preload("Peminjaman").First(&anggota, id).Error
    return anggota, err
}

func (r *anggotaRepository) FindByJudul(keyword string) ([]model.Anggota, error) {
    var anggotaList []model.Anggota
    err := r.db.Preload("Peminjaman").
        Where("nama LIKE ?", "%"+keyword+"%").
        Find(&anggotaList).Error
    return anggotaList, err
}

func (r *anggotaRepository) Create(anggota model.Anggota) (model.Anggota, error) {
    err := r.db.Create(&anggota).Error
    return anggota, err
}

func (r *anggotaRepository) Update(anggota model.Anggota, input model.Anggota) (model.Anggota, error) {
    err := r.db.Model(&anggota).Updates(input).Error
    return anggota, err
}

func (r *anggotaRepository) Delete(anggota model.Anggota) error {
    return r.db.Delete(&anggota).Error
}