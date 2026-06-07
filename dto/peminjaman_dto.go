package dto

// Request
type PeminjamanRequest struct {
	AnggotaID     uint                  `json:"anggota_id" binding:"required"`
	PetugasID     uint                  `json:"petugas_id" binding:"required"`
	TglPinjam     string                `json:"tgl_pinjam" binding:"required"`
	TglJatuhTempo string                `json:"tgl_jatuh_tempo" binding:"required"`
	DetailBuku    []DetailPeminjamanRequest `json:"detail_buku" binding:"required"`
}

type DetailPeminjamanRequest struct {
	BukuID uint `json:"buku_id" binding:"required"`
	Jumlah int  `json:"jumlah" binding:"required,gt=0"`
}

// Response
type PeminjamanResponse struct {
	ID            uint                    `json:"id"`
	TglPinjam     string                  `json:"tgl_pinjam"`
	TglJatuhTempo string                  `json:"tgl_jatuh_tempo"`
	Status        string                  `json:"status"`
	Anggota       AnggotaResponse         `json:"anggota"`
	Petugas       PetugasResponse         `json:"petugas"`
	Detail        []DetailPeminjamanResponse `json:"detail_peminjaman"`
}

type DetailPeminjamanResponse struct {
	ID     uint         `json:"id"`
	Jumlah int          `json:"jumlah"`
	Buku   BukuResponse `json:"buku"`
}