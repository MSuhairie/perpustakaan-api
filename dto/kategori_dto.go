package dto

// Request
type KategoriRequest struct {
	NamaKategori string `json:"nama_kategori" binding:"required"`
}

// Response
type KategoriResponse struct {
	ID           uint   `json:"id"`
	NamaKategori string `json:"nama_kategori"`
}