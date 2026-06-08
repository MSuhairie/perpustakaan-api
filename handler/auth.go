package handler

import (
	"net/http"
	"perpustakaan-api/database"
	"perpustakaan-api/helper"
	"perpustakaan-api/middleware"
	"perpustakaan-api/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Nama     string `json:"nama" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required,min=6"`
	Role     string `json:"role" binding:"required"`
}

type LoginInput struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
        helper.ResponseValidationError(c, helper.PesanError(err))
		return
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
        helper.ResponseInternalError(c, "Gagal proses password")
		return
	}

	petugas := model.Petugas{
		Nama:     input.Nama,
		Username: input.Username,
		Password: string(hashedPassword),
		Role:     input.Role,
	}

	if err := database.DB.Create(&petugas).Error; err != nil {
		helper.ResponseBadRequest(c, "Username sudah digunakan")
		return
	}

	helper.ResponseCreated(c, "Registrasi berhasil", nil)
}

func Login(c *gin.Context) {
	var input LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}

	// Cari petugas by username
	var petugas model.Petugas
	if err := database.DB.Where("username = ?", input.Username).First(&petugas).Error; err != nil {
		helper.ResponseUnauthorized(c, "Username atau password salah")
		return
	}

	// Cek password
	if err := bcrypt.CompareHashAndPassword([]byte(petugas.Password), []byte(input.Password)); err != nil {
		helper.ResponseUnauthorized(c, "Username atau password salah")
		return
	}

	// Buat token JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"petugas_id": petugas.ID,
		"username":   petugas.Username,
		"exp":        time.Now().Add(24 * time.Hour).Unix(),
	})

	tokenString, err := token.SignedString(middleware.GetSecretKey())
	if err != nil {
		helper.ResponseInternalError(c, "Gagal buat token")		
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"token":   tokenString,
		"petugas": gin.H{
			"id"       : petugas.ID,
			"nama"     : petugas.Nama,
			"username" : petugas.Username,
			"role"     : petugas.Role,
		},
	})
}
