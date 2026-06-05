package usecase

import (
    "errors"
    "perpustakaan-api/model"
    "perpustakaan-api/repository"
)

type RakUsecase interface {
    GetAllRak() ([]model.Rak, error)
    GetRakByID(id string) (model.Rak, error)
    SearchRak(keyword string) ([]model.Rak, error)
    CreateRak(rak model.Rak) (model.Rak, error)
    UpdateRak(id string, input model.Rak) (model.Rak, error)
    DeleteRak(id string) error
}

type rakUsecase struct {
    repo repository.RakRepository
}

func NewRakUsecase(repo repository.RakRepository) RakUsecase {
    return &rakUsecase{repo}
}

func (u *rakUsecase) GetAllRak() ([]model.Rak, error) {
    return u.repo.FindAll()
}

func (u *rakUsecase) GetRakByID(id string) (model.Rak, error) {
    rak, err := u.repo.FindByID(id)
    if err != nil {
        return rak, errors.New("rak tidak ditemukan")
    }
    return rak, nil
}

func (u *rakUsecase) SearchRak(keyword string) ([]model.Rak, error) {
    if keyword == "" {
        return nil, errors.New("keyword tidak boleh kosong")
    }
    return u.repo.FindByJudul(keyword)
}

func (u *rakUsecase) CreateRak(rak model.Rak) (model.Rak, error) {
    if rak.KodeRak == "" {
        return rak, errors.New("kode rak tidak boleh kosong")
    }
    if rak.Lokasi == "" {
        return rak, errors.New("lokasi tidak boleh kosong")
    }
    existingRak, _ := u.repo.FindByKodeRak(rak.KodeRak)
    if existingRak.ID != 0 {
        return rak, errors.New("kode rak sudah digunakan")
    }
    return u.repo.Create(rak)
}

func (u *rakUsecase) UpdateRak(id string, input model.Rak) (model.Rak, error) {
    rak, err := u.repo.FindByID(id)
    if err != nil {
        return rak, errors.New("rak tidak ditemukan")
    }
    return u.repo.Update(rak, input)
}

func (u *rakUsecase) DeleteRak(id string) error {
    rak, err := u.repo.FindByID(id)
    if err != nil {
        return errors.New("rak tidak ditemukan")
    }
    return u.repo.Delete(rak)
}