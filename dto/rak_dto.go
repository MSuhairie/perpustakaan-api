package dto

// Request
type RakRequest struct {
	KodeRak string `json:"kode_rak" binding:"required"`
	Lokasi  string `json:"lokasi" binding:"required"`
}

// Response
type RakResponse struct {
	ID      uint   `json:"id"`
	KodeRak string `json:"kode_rak"`
	Lokasi  string `json:"lokasi"`
}