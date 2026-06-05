package repository

import (
	"perpustakaan-api/model"

	"gorm.io/gorm"
)

type KategoriRepository interface {
	FindAll() ([]model.Kategori, error)
	FindByID(id string) (model.Kategori, error)
	FindByJudul(keyword string) ([]model.Kategori, error)
	Create(kategori model.Kategori) (model.Kategori, error)
	Update(kategori model.Kategori, input model.Kategori) (model.Kategori, error)
	Delete(kategori model.Kategori) error
}

type kategoriRepository struct {
	db *gorm.DB
}

func NewKategoriRepository(db *gorm.DB) KategoriRepository {
	return &kategoriRepository{db}
}

func (r *kategoriRepository) FindAll() ([]model.Kategori, error) {
	var kategoriList []model.Kategori
	err := r.db.Preload("Buku").Find(&kategoriList).Error
	return kategoriList, err
}

func (r *kategoriRepository) FindByID(id string) (model.Kategori, error) {
    var kategori model.Kategori
    err := r.db.Preload("Buku").First(&kategori, id).Error
    return kategori, err
}

func (r *kategoriRepository) FindByJudul(keyword string) ([]model.Kategori, error) {
    var kategoriList []model.Kategori
    err := r.db.Preload("Buku").
        Where("nama_kategori LIKE ?", "%"+keyword+"%").
        Find(&kategoriList).Error
    return kategoriList, err
}

func (r *kategoriRepository) Create(kategori model.Kategori) (model.Kategori, error) {
    err := r.db.Create(&kategori).Error
    return kategori, err
}

func (r *kategoriRepository) Update(kategori model.Kategori, input model.Kategori) (model.Kategori, error) {
    err := r.db.Model(&kategori).Updates(input).Error
    return kategori, err
}

func (r *kategoriRepository) Delete(kategori model.Kategori) error {
    return r.db.Delete(&kategori).Error
}