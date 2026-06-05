package usecase

import (
    "errors"
    "perpustakaan-api/model"
    "perpustakaan-api/repository"
)

type BukuUsecase interface {
    GetAllBuku() ([]model.Buku, error)
    GetBukuByID(id string) (model.Buku, error)
    SearchBuku(keyword string) ([]model.Buku, error)
    CreateBuku(buku model.Buku) (model.Buku, error)
    UpdateBuku(id string, input model.Buku) (model.Buku, error)
    DeleteBuku(id string) error
}

type bukuUsecase struct {
    repo repository.BukuRepository
}

func NewBukuUsecase(repo repository.BukuRepository) BukuUsecase {
    return &bukuUsecase{repo}
}

func (u *bukuUsecase) GetAllBuku() ([]model.Buku, error) {
    return u.repo.FindAll()
}

func (u *bukuUsecase) GetBukuByID(id string) (model.Buku, error) {
    buku, err := u.repo.FindByID(id)
    if err != nil {
        return buku, errors.New("buku tidak ditemukan")
    }
    return buku, nil
}

func (u *bukuUsecase) SearchBuku(keyword string) ([]model.Buku, error) {
    if keyword == "" {
        return nil, errors.New("keyword tidak boleh kosong")
    }
    return u.repo.FindByJudul(keyword)
}

func (u *bukuUsecase) CreateBuku(buku model.Buku) (model.Buku, error) {
    if buku.Judul == "" {
        return buku, errors.New("judul buku tidak boleh kosong")
    }
    if buku.ISBN == "" {
        return buku, errors.New("ISBN tidak boleh kosong")
    }
    if buku.Stok < 0 {
        return buku, errors.New("stok tidak boleh minus")
    }
    return u.repo.Create(buku)
}

func (u *bukuUsecase) UpdateBuku(id string, input model.Buku) (model.Buku, error) {
    buku, err := u.repo.FindByID(id)
    if err != nil {
        return buku, errors.New("buku tidak ditemukan")
    }
    return u.repo.Update(buku, input)
}

func (u *bukuUsecase) DeleteBuku(id string) error {
    buku, err := u.repo.FindByID(id)
    if err != nil {
        return errors.New("buku tidak ditemukan")
    }
    return u.repo.Delete(buku)
}