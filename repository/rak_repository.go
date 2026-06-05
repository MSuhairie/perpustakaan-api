package repository

import (
	"perpustakaan-api/model"

	"gorm.io/gorm"
)

type RakRepository interface {
	FindAll() ([]model.Rak, error)
	FindByID(id string) (model.Rak, error)
	FindByKodeRak(koderak string) (model.Rak, error)
	FindByJudul(keyword string) ([]model.Rak, error)
	Create(rak model.Rak) (model.Rak, error)
	Update(rak model.Rak, input model.Rak) (model.Rak, error)
	Delete(rak model.Rak) error
}

type rakRepository struct {
	db *gorm.DB
}

func NewRakRepository(db *gorm.DB) RakRepository {
	return &rakRepository{db}
}

func (r *rakRepository) FindAll() ([]model.Rak, error) {
	var rakList []model.Rak
	err := r.db.Preload("Buku").Find(&rakList).Error
	return rakList, err
}

func (r *rakRepository) FindByID(id string) (model.Rak, error) {
    var rak model.Rak
    err := r.db.Preload("Buku").First(&rak, id).Error
    return rak, err
}

func (r *rakRepository) FindByKodeRak(koderak string) (model.Rak, error) {
    var rak model.Rak
    err := r.db.Where("kode_rak = ?", koderak).First(&rak).Error
    return rak, err
}

func (r *rakRepository) FindByJudul(keyword string) ([]model.Rak, error) {
    var rakList []model.Rak
    err := r.db.Preload("Buku").
        Where("kode_rak LIKE ?", "%"+keyword+"%").
        Find(&rakList).Error
    return rakList, err
}

func (r *rakRepository) Create(rak model.Rak) (model.Rak, error) {
    err := r.db.Create(&rak).Error
    return rak, err
}

func (r *rakRepository) Update(rak model.Rak, input model.Rak) (model.Rak, error) {
    err := r.db.Model(&rak).Updates(input).Error
    return rak, err
}

func (r *rakRepository) Delete(rak model.Rak) error {
    return r.db.Delete(&rak).Error
}