package usecase

import (
    "errors"
    "perpustakaan-api/model"
    "perpustakaan-api/repository"
)

type AnggotaUsecase interface {
    GetAllAnggota() ([]model.Anggota, error)
    GetAnggotaByID(id string) (model.Anggota, error)
    SearchAnggota(keyword string) ([]model.Anggota, error)
    CreateAnggota(anggota model.Anggota) (model.Anggota, error)
    UpdateAnggota(id string, input model.Anggota) (model.Anggota, error)
    DeleteAnggota(id string) error
}

type anggotaUsecase struct {
    repo repository.AnggotaRepository
}

func NewAnggotaUsecase(repo repository.AnggotaRepository) AnggotaUsecase {
    return &anggotaUsecase{repo}
}

func (u *anggotaUsecase) GetAllAnggota() ([]model.Anggota, error) {
    return u.repo.FindAll()
}

func (u *anggotaUsecase) GetAnggotaByID(id string) (model.Anggota, error) {
    anggota, err := u.repo.FindByID(id)
    if err != nil {
        return anggota, errors.New("anggota tidak ditemukan")
    }
    return anggota, nil
}

func (u *anggotaUsecase) SearchAnggota(keyword string) ([]model.Anggota, error) {
    if keyword == "" {
        return nil, errors.New("keyword tidak boleh kosong")
    }
    return u.repo.FindByJudul(keyword)
}

func (u *anggotaUsecase) CreateAnggota(anggota model.Anggota) (model.Anggota, error) {
    if anggota.Nama == "" {
        return anggota, errors.New("nama anggota tidak boleh kosong")
    }
    if anggota.Alamat == "" {
        return anggota, errors.New("alamat tidak boleh kosong")
    }
    if anggota.NoHP == "" {
        return anggota, errors.New("nohp tidak boleh kosong")
    }
    if anggota.TglDaftar == "" {
        return anggota, errors.New("tgldaftar tidak boleh kosong")
    }
    if anggota.Status == "" {
        return anggota, errors.New("status tidak boleh kosong")
    }
    return u.repo.Create(anggota)
}

func (u *anggotaUsecase) UpdateAnggota(id string, input model.Anggota) (model.Anggota, error) {
    anggota, err := u.repo.FindByID(id)
    if err != nil {
        return anggota, errors.New("anggota tidak ditemukan")
    }
    _, err = u.repo.Update(anggota, input)
    if err != nil {
        return anggota, err
    }
    return u.repo.FindByID(id)
}

func (u *anggotaUsecase) DeleteAnggota(id string) error {
    anggota, err := u.repo.FindByID(id)
    if err != nil {
        return errors.New("anggota tidak ditemukan")
    }
    return u.repo.Delete(anggota)
}