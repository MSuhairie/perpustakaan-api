package dto

// Request
type PengembalianRequest struct {
	PeminjamanID uint                     `json:"peminjaman_id" binding:"required"`
	PetugasID    uint                     `json:"petugas_id" binding:"required"`
	TglKembali   string                   `json:"tgl_kembali" binding:"required"`
	Detail       []DetailPengembalianRequest `json:"detail" binding:"required"`
}

type DetailPengembalianRequest struct {
	DetailPeminjamanID uint `json:"detail_peminjaman_id" binding:"required"`
	KondisiBuku        int  `json:"kondisi_buku" binding:"required"`
}

// Response
type DendaResponse struct {
	ID          uint   `json:"id"`
	JumlahDenda int    `json:"jumlah_denda"`
	Keterangan  string `json:"keterangan"`
}

type DetailPengembalianResponse struct {
	ID               uint                     `json:"id"`
	KondisiBuku      int                      `json:"kondisi_buku"`
	DetailPeminjaman DetailPeminjamanResponse  `json:"detail_peminjaman"`
}

type PengembalianResponse struct {
	ID                 uint                         `json:"id"`
	TglKembali         string                       `json:"tgl_kembali"`
	Peminjaman         PeminjamanResponse           `json:"peminjaman"`
	Petugas            PetugasResponse              `json:"petugas"`
	DetailPengembalian []DetailPengembalianResponse `json:"detail_pengembalian"`
	Denda              *DendaResponse               `json:"denda,omitempty"`
}