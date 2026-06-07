package usecase

import (
    "errors"
    "perpustakaan-api/model"
    "perpustakaan-api/repository"
)

type KategoriUsecase interface {
    GetAllKategori() ([]model.Kategori, error)
    GetKategoriByID(id string) (model.Kategori, error)
    SearchKategori(keyword string) ([]model.Kategori, error)
    CreateKategori(kategori model.Kategori) (model.Kategori, error)
    UpdateKategori(id string, input model.Kategori) (model.Kategori, error)
    DeleteKategori(id string) error
}

type kategoriUsecase struct {
    repo repository.KategoriRepository
}

func NewKategoriUsecase(repo repository.KategoriRepository) KategoriUsecase {
    return &kategoriUsecase{repo}
}

func (u *kategoriUsecase) GetAllKategori() ([]model.Kategori, error) {
    return u.repo.FindAll()
}

func (u *kategoriUsecase) GetKategoriByID(id string) (model.Kategori, error) {
    kategori, err := u.repo.FindByID(id)
    if err != nil {
        return kategori, errors.New("kategori tidak ditemukan")
    }
    return kategori, nil
}

func (u *kategoriUsecase) SearchKategori(keyword string) ([]model.Kategori, error) {
    if keyword == "" {
        return nil, errors.New("keyword tidak boleh kosong")
    }
    return u.repo.FindByJudul(keyword)
}

func (u *kategoriUsecase) CreateKategori(kategori model.Kategori) (model.Kategori, error) {
    if kategori.NamaKategori == "" {
        return kategori, errors.New("nama kategori tidak boleh kosong")
    }
    return u.repo.Create(kategori)
}

func (u *kategoriUsecase) UpdateKategori(id string, input model.Kategori) (model.Kategori, error) {
    kategori, err := u.repo.FindByID(id)
    if err != nil {
        return kategori, errors.New("kategori tidak ditemukan")
    }
    _, err = u.repo.Update(kategori, input)
    if err != nil {
        return kategori, err
    }
    return u.repo.FindByID(id)
}

func (u *kategoriUsecase) DeleteKategori(id string) error {
    kategori, err := u.repo.FindByID(id)
    if err != nil {
        return errors.New("kategori tidak ditemukan")
    }
    return u.repo.Delete(kategori)
}