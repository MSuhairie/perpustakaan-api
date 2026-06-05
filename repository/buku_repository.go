package repository

import (
	"perpustakaan-api/model"

    "gorm.io/gorm"
)

type BukuRepository interface {
	FindAll() ([]model.Buku, error)
	FindByID(id string) (model.Buku, error)
    FindByJudul(keyword string) ([]model.Buku, error)
    Create(buku model.Buku) (model.Buku, error)
    Update(buku model.Buku, input model.Buku) (model.Buku, error)
    Delete(buku model.Buku) error
}

type bukuRepository struct {
	db *gorm.DB
}

func NewBukuRepository(db *gorm.DB) BukuRepository {
	return &bukuRepository{db}
}

func (r *bukuRepository) FindAll() ([]model.Buku, error) {
	var bukuList []model.Buku
	err := r.db.Preload("Kategori").Preload("Rak").Find(&bukuList).Error
	return bukuList, err
}

func (r *bukuRepository) FindByID(id string) (model.Buku, error) {
    var buku model.Buku
    err := r.db.Preload("Kategori").Preload("Rak").First(&buku, id).Error
    return buku, err
}

func (r *bukuRepository) FindByJudul(keyword string) ([]model.Buku, error) {
    var bukuList []model.Buku
    err := r.db.Preload("Kategori").Preload("Rak").
        Where("judul LIKE ?", "%"+keyword+"%").
        Find(&bukuList).Error
    return bukuList, err
}

func (r *bukuRepository) Create(buku model.Buku) (model.Buku, error) {
    err := r.db.Create(&buku).Error
    return buku, err
}

func (r *bukuRepository) Update(buku model.Buku, input model.Buku) (model.Buku, error) {
    err := r.db.Model(&buku).Updates(input).Error
    return buku, err
}

func (r *bukuRepository) Delete(buku model.Buku) error {
    return r.db.Delete(&buku).Error
}