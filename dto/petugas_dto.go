package dto

// Request Register
type RegisterRequest struct {
	Nama     string `json:"nama" binding:"required"`
	Username string `json:"username" binding:"required,min=3"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"`
}

// Request Login
type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Response Petugas
type PetugasResponse struct {
	ID       uint   `json:"id"`
	Nama     string `json:"nama"`
	Username string `json:"username"`
	Role     string `json:"role"`
}

// Response Login
type LoginResponse struct {
	Token   string          `json:"token"`
	Petugas PetugasResponse `json:"petugas"`
}