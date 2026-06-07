package dto

// Request
type AnggotaRequest struct {
	Nama      string `json:"nama" binding:"required,min=3"`
	Alamat    string `json:"alamat" binding:"required"`
	NoHP      string `json:"no_hp" binding:"required,min=10,max=13"`
	TglDaftar string `json:"tgl_daftar" binding:"required"`
	Status    string `json:"status" binding:"required"`
}

// Response
type AnggotaResponse struct {
	ID        uint   `json:"id"`
	Nama      string `json:"nama"`
	Alamat    string `json:"alamat"`
	NoHP      string `json:"no_hp"`
	TglDaftar string `json:"tgl_daftar"`
	Status    string `json:"status"`
}