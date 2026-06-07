package dto

// Request — untuk Create & Update
type BukuRequest struct {
	KategoriID  uint   `json:"kategori_id" binding:"required"`
	RakID       uint   `json:"rak_id" binding:"required"`
	Judul       string `json:"judul" binding:"required"`
	Penulis     string `json:"penulis" binding:"required"`
	Penerbit    string `json:"penerbit" binding:"required"`
	TahunTerbit string `json:"tahun_terbit" binding:"required"`
	ISBN        string `json:"isbn" binding:"required"`
	Stok        int    `json:"stok" binding:"gte=0"`
}

// Response — data yang dikembalikan ke client
type BukuResponse struct {
	ID          uint             `json:"id"`
	Judul       string           `json:"judul"`
	Penulis     string           `json:"penulis"`
	Penerbit    string           `json:"penerbit"`
	TahunTerbit string           `json:"tahun_terbit"`
	ISBN        string           `json:"isbn"`
	Stok        int              `json:"stok"`
	Foto        string           `json:"foto,omitempty"`
	Kategori    KategoriResponse `json:"kategori,omitempty"`
	Rak         RakResponse      `json:"rak,omitempty"`
}