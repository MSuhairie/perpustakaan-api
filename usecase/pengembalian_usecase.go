package usecase

import (
    "errors"
    "fmt"
    "perpustakaan-api/model"
    "perpustakaan-api/repository"
    "time"
)

type PengembalianResponse struct {
    model.Pengembalian
    Denda *model.Denda `json:"denda,omitempty"`
}

type PengembalianUsecase interface {
    GetAll() ([]model.Pengembalian, error)
    GetByID(id string) (model.Pengembalian, error)
    Create(input CreatePengembalianInput) (PengembalianResponse, error)
}

type CreatePengembalianInput struct {
    PeminjamanID uint                     `json:"peminjaman_id" binding:"required"`
    PetugasID    uint                     `json:"petugas_id" binding:"required"`
    TglKembali   string                   `json:"tgl_kembali" binding:"required"`
    Detail       []DetailPengembalianInput `json:"detail" binding:"required"`
}

type DetailPengembalianInput struct {
    DetailPeminjamanID uint `json:"detail_peminjaman_id" binding:"required"`
    KondisiBuku        int  `json:"kondisi_buku" binding:"required"`
    // kondisi: 1=baik, 2=rusak ringan, 3=rusak berat
}

type pengembalianUsecase struct {
    repo         repository.PengembalianRepository
    peminjamanRepo repository.PeminjamanRepository
    bukuRepo     repository.BukuRepository
    dendaRepo    repository.DendaRepository
}

func NewPengembalianUsecase(
    repo repository.PengembalianRepository,
    peminjamanRepo repository.PeminjamanRepository,
    bukuRepo repository.BukuRepository,
    dendaRepo repository.DendaRepository,
) PengembalianUsecase {
    return &pengembalianUsecase{repo, peminjamanRepo, bukuRepo, dendaRepo}
}

func (u *pengembalianUsecase) GetAll() ([]model.Pengembalian, error) {
    return u.repo.FindAll()
}

func (u *pengembalianUsecase) GetByID(id string) (model.Pengembalian, error) {
    p, err := u.repo.FindByID(id)
    if err != nil {
        return p, errors.New("pengembalian tidak ditemukan")
    }
    return p, nil
}

func (u *pengembalianUsecase) Create(input CreatePengembalianInput) (PengembalianResponse, error) {
    // 1. Cari data peminjaman
    peminjaman, err := u.peminjamanRepo.FindByID(fmt.Sprintf("%d", input.PeminjamanID))
    if err != nil {
        return PengembalianResponse{}, errors.New("peminjaman tidak ditemukan")
    }

    // 2. Cek status peminjaman
    if peminjaman.Status == "dikembalikan" {
        return PengembalianResponse{}, errors.New("buku sudah dikembalikan sebelumnya")
    }

    // 3. Hitung denda kalau terlambat
    tglJatuhTempo, _ := time.Parse("2006-01-02", peminjaman.TglJatuhTempo)
    tglKembali, _    := time.Parse("2006-01-02", input.TglKembali)

    var dendaResult *model.Denda
    if tglKembali.After(tglJatuhTempo) {
        selisihHari := int(tglKembali.Sub(tglJatuhTempo).Hours() / 24)
        jumlahBuku  := len(peminjaman.DetailPeminjaman)
        jumlahDenda := selisihHari * 1000 * jumlahBuku

        dendaBaru := model.Denda{
            PeminjamanID: input.PeminjamanID,
            JumlahDenda:  jumlahDenda,
            Keterangan: fmt.Sprintf(
                "Terlambat %d hari, denda Rp %d",
                selisihHari,
                jumlahDenda,
            ),
        }
        denda, err := u.dendaRepo.Create(dendaBaru)
        if err == nil {
            dendaResult = &denda  // ← simpan hasil denda
        }
    }

    // 4. Buat detail pengembalian
    var detailList []model.DetailPengembalian
    for _, d := range input.Detail {
        detailList = append(detailList, model.DetailPengembalian{
            DetailPeminjamanID: d.DetailPeminjamanID,
            KondisiBuku:        d.KondisiBuku,
        })
    }

    // 5. Buat pengembalian
    pengembalian := model.Pengembalian{
        PeminjamanID:       input.PeminjamanID,
        PetugasID:          input.PetugasID,
        TglKembali:         input.TglKembali,
        DetailPengembalian: detailList,
    }

    result, err := u.repo.Create(pengembalian)
    if err != nil {
        return PengembalianResponse{}, err
    }

    // 6. Kembalikan stok buku
    for _, detail := range peminjaman.DetailPeminjaman {
        buku, err := u.bukuRepo.FindByID(fmt.Sprintf("%d", detail.BukuID))
        if err == nil {
            u.bukuRepo.Update(buku, model.Buku{Stok: buku.Stok + detail.Jumlah})
        }
    }

    // 7. Update status peminjaman jadi "dikembalikan"
    if tglKembali.After(tglJatuhTempo) {
        u.peminjamanRepo.Update(peminjaman, model.Peminjaman{Status: "terlambat"})
    } else {
        u.peminjamanRepo.Update(peminjaman, model.Peminjaman{Status: "dikembalikan"})
    }

    // 8 — Fetch ulang dengan semua relasi
    finalResult, err := u.repo.FindByID(fmt.Sprintf("%d", result.ID))
    if err != nil {
        return PengembalianResponse{}, err
    }

    // Gabungkan pengembalian + denda
    return PengembalianResponse{
        Pengembalian: finalResult,
        Denda:        dendaResult,
    }, nil
}