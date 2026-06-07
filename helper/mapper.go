package helper

import (
	"perpustakaan-api/dto"
	"perpustakaan-api/model"
)

// Kategori
func ToKategoriResponse(k model.Kategori) dto.KategoriResponse {
	return dto.KategoriResponse{
		ID:           k.ID,
		NamaKategori: k.NamaKategori,
	}
}

func ToKategoriResponseList(kategoriList []model.Kategori) []dto.KategoriResponse {
	var result []dto.KategoriResponse
	for _, k := range kategoriList {
		result = append(result, ToKategoriResponse(k))
	}
	return result
}

// Rak
func ToRakResponse(r model.Rak) dto.RakResponse {
	return dto.RakResponse{
		ID:      r.ID,
		KodeRak: r.KodeRak,
		Lokasi:  r.Lokasi,
	}
}

func ToRakResponseList(rakList []model.Rak) []dto.RakResponse {
	var result []dto.RakResponse
	for _, r := range rakList {
		result = append(result, ToRakResponse(r))
	}
	return result
}

// Buku
func ToBukuResponse(b model.Buku) dto.BukuResponse {
	return dto.BukuResponse{
		ID:          b.ID,
		Judul:       b.Judul,
		Penulis:     b.Penulis,
		Penerbit:    b.Penerbit,
		TahunTerbit: b.TahunTerbit,
		ISBN:        b.ISBN,
		Stok:        b.Stok,
		Foto:        b.Foto,
		Kategori:    ToKategoriResponse(b.Kategori),
		Rak:         ToRakResponse(b.Rak),
	}
}

func ToBukuResponseList(bukuList []model.Buku) []dto.BukuResponse {
	var result []dto.BukuResponse
	for _, b := range bukuList {
		result = append(result, ToBukuResponse(b))
	}
	return result
}

// Anggota
func ToAnggotaResponse(a model.Anggota) dto.AnggotaResponse {
	return dto.AnggotaResponse{
		ID:        a.ID,
		Nama:      a.Nama,
		Alamat:    a.Alamat,
		NoHP:      a.NoHP,
		TglDaftar: a.TglDaftar,
		Status:    a.Status,
	}
}

func ToAnggotaResponseList(list []model.Anggota) []dto.AnggotaResponse {
	var result []dto.AnggotaResponse
	for _, a := range list {
		result = append(result, ToAnggotaResponse(a))
	}
	return result
}

// Petugas
func ToPetugasResponse(p model.Petugas) dto.PetugasResponse {
	return dto.PetugasResponse{
		ID:       p.ID,
		Nama:     p.Nama,
		Username: p.Username,
		Role:     p.Role,
	}
}

// Detail Peminjaman
func ToDetailPeminjamanResponse(d model.DetailPeminjaman) dto.DetailPeminjamanResponse {
	return dto.DetailPeminjamanResponse{
		ID:     d.ID,
		Jumlah: d.Jumlah,
		Buku:   ToBukuResponse(d.Buku),
	}
}

// Peminjaman
func ToPeminjamanResponse(p model.Peminjaman) dto.PeminjamanResponse {
	var details []dto.DetailPeminjamanResponse
	for _, d := range p.DetailPeminjaman {
		details = append(details, ToDetailPeminjamanResponse(d))
	}
	return dto.PeminjamanResponse{
		ID:            p.ID,
		TglPinjam:     p.TglPinjam,
		TglJatuhTempo: p.TglJatuhTempo,
		Status:        p.Status,
		Anggota:       ToAnggotaResponse(p.Anggota),
		Petugas:       ToPetugasResponse(p.Petugas),
		Detail:        details,
	}
}

func ToPeminjamanResponseList(list []model.Peminjaman) []dto.PeminjamanResponse {
	var result []dto.PeminjamanResponse
	for _, p := range list {
		result = append(result, ToPeminjamanResponse(p))
	}
	return result
}

// Detail Pengembalian
func ToDetailPengembalianResponse(d model.DetailPengembalian) dto.DetailPengembalianResponse {
	return dto.DetailPengembalianResponse{
		ID:          d.ID,
		KondisiBuku: d.KondisiBuku,
		DetailPeminjaman: ToDetailPeminjamanResponse(d.DetailPeminjaman),
	}
}

// Pengembalian
func ToPengembalianResponse(p model.Pengembalian, denda *model.Denda) dto.PengembalianResponse {
	var details []dto.DetailPengembalianResponse
	for _, d := range p.DetailPengembalian {
		details = append(details, ToDetailPengembalianResponse(d))
	}

	response := dto.PengembalianResponse{
		ID:                 p.ID,
		TglKembali:         p.TglKembali,
		Peminjaman:         ToPeminjamanResponse(p.Peminjaman),
		Petugas:            ToPetugasResponse(p.Petugas),
		DetailPengembalian: details,
	}

	// Attach denda kalau ada
	if denda != nil {
		response.Denda = &dto.DendaResponse{
			ID:          denda.ID,
			JumlahDenda: denda.JumlahDenda,
			Keterangan:  denda.Keterangan,
		}
	}

	return response
}