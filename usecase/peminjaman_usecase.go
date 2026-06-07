package usecase

import (
	"errors"
	"fmt"
	"perpustakaan-api/model"
	"perpustakaan-api/repository"
)

type PeminjamanUsecase interface {
    GetAll() ([]model.Peminjaman, error)
    GetByID(id string) (model.Peminjaman, error)
    GetByAnggota(anggotaID string) ([]model.Peminjaman, error)
    Create(input CreatePeminjamanInput) (model.Peminjaman, error)
    UpdateStatus(id string, status string) (model.Peminjaman, error)
    Delete(id string) error
}

// Input struct untuk create peminjaman
type CreatePeminjamanInput struct {
    AnggotaID     uint                  `json:"anggota_id" binding:"required"`
    PetugasID     uint                  `json:"petugas_id" binding:"required"`
    TglPinjam     string                `json:"tgl_pinjam" binding:"required"`
    TglJatuhTempo string                `json:"tgl_jatuh_tempo" binding:"required"`
    DetailBuku    []DetailPeminjamanInput `json:"detail_buku" binding:"required"`
}

type DetailPeminjamanInput struct {
    BukuID uint `json:"buku_id" binding:"required"`
    Jumlah int  `json:"jumlah" binding:"required,gt=0"`
}

type peminjamanUsecase struct {
    repo     repository.PeminjamanRepository
    bukuRepo repository.BukuRepository
}

func NewPeminjamanUsecase(
    repo repository.PeminjamanRepository,
    bukuRepo repository.BukuRepository,
) PeminjamanUsecase {
    return &peminjamanUsecase{repo, bukuRepo}
}

func (u *peminjamanUsecase) GetAll() ([]model.Peminjaman, error) {
    return u.repo.FindAll()
}

func (u *peminjamanUsecase) GetByID(id string) (model.Peminjaman, error) {
    p, err := u.repo.FindByID(id)
    if err != nil {
        return p, errors.New("peminjaman tidak ditemukan")
    }
    return p, nil
}

func (u *peminjamanUsecase) GetByAnggota(anggotaID string) ([]model.Peminjaman, error) {
    return u.repo.FindByAnggota(anggotaID)
}

func (u *peminjamanUsecase) Create(input CreatePeminjamanInput) (model.Peminjaman, error) {
    // Validasi detail buku tidak boleh kosong
    if len(input.DetailBuku) == 0 {
        return model.Peminjaman{}, errors.New("minimal 1 buku harus dipilih")
    }

    // Cek stok setiap buku & kurangi stok
    for _, detail := range input.DetailBuku {
        buku, err := u.bukuRepo.FindByID(fmt.Sprintf("%d", detail.BukuID))
        if err != nil {
            return model.Peminjaman{}, errors.New("buku tidak ditemukan")
        }
        if buku.Stok < detail.Jumlah {
            return model.Peminjaman{}, fmt.Errorf("stok buku '%s' tidak mencukupi, stok tersedia: %d", buku.Judul, buku.Stok)
        }
        // Kurangi stok
        u.bukuRepo.Update(buku, model.Buku{Stok: buku.Stok - detail.Jumlah})
    }

    // Buat detail peminjaman
    var detailList []model.DetailPeminjaman
    for _, d := range input.DetailBuku {
        detailList = append(detailList, model.DetailPeminjaman{
            BukuID: d.BukuID,
            Jumlah: d.Jumlah,
        })
    }

    // Buat peminjaman
    peminjaman := model.Peminjaman{
        AnggotaID:        input.AnggotaID,
        PetugasID:        input.PetugasID,
        TglPinjam:        input.TglPinjam,
        TglJatuhTempo:    input.TglJatuhTempo,
        Status:           "dipinjam",
        DetailPeminjaman: detailList,
    }

    return u.repo.Create(peminjaman)
}

func (u *peminjamanUsecase) UpdateStatus(id string, status string) (model.Peminjaman, error) {
    peminjaman, err := u.repo.FindByID(id)
    if err != nil {
        return peminjaman, errors.New("peminjaman tidak ditemukan")
    }

    // Validasi status
    validStatus := map[string]bool{
        "dipinjam":    true,
        "dikembalikan": true,
        "terlambat":   true,
    }
    if !validStatus[status] {
        return peminjaman, errors.New("status tidak valid")
    }

    return u.repo.Update(peminjaman, model.Peminjaman{Status: status})
}

func (u *peminjamanUsecase) Delete(id string) error {
    peminjaman, err := u.repo.FindByID(id)
    if err != nil {
        return errors.New("peminjaman tidak ditemukan")
    }

    // Kembalikan stok buku
    for _, detail := range peminjaman.DetailPeminjaman {
        buku, err := u.bukuRepo.FindByID(fmt.Sprintf("%d", detail.BukuID))
        if err == nil {
            u.bukuRepo.Update(buku, model.Buku{Stok: buku.Stok + detail.Jumlah})
        }
    }

    return u.repo.Delete(peminjaman)
}