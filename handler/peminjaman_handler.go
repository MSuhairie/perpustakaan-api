package handler

import (
	"perpustakaan-api/helper"
	"perpustakaan-api/usecase"

	"github.com/gin-gonic/gin"
)

type PeminjamanHandler struct {
	usecase usecase.PeminjamanUsecase
}

func NewPeminjamanHandler(uc usecase.PeminjamanUsecase) *PeminjamanHandler {
	return &PeminjamanHandler{uc}
}

func (h *PeminjamanHandler) GetAll(c *gin.Context) {
	// Cek apakah ada query parameter anggota_id
    anggotaID := c.Query("anggota_id")

    if anggotaID != "" {
        // Filter by anggota
        list, err := h.usecase.GetByAnggota(anggotaID)
        if err != nil {
            helper.ResponseInternalError(c, err.Error())
            return
        }
        helper.ResponseOK(c, "Data peminjaman anggota berhasil diambil", list)
        return
    }

	list, err := h.usecase.GetAll()
	if err != nil {
		helper.ResponseInternalError(c, err.Error())
		return
	}
	helper.ResponseOK(c, "Data peminjaman berhasil diambil", list)
}

func (h *PeminjamanHandler) GetByID(c *gin.Context) {
	p, err := h.usecase.GetByID(c.Param("id"))
	if err != nil {
		helper.ResponseNotFound(c, err.Error())
		return
	}
	helper.ResponseOK(c, "Data peminjaman berhasil diambil", p)
}

func (h *PeminjamanHandler) Create(c *gin.Context) {
	var input usecase.CreatePeminjamanInput
	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ResponseValidationError(c, helper.PesanError(err))
		return
	}
	result, err := h.usecase.Create(input)
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}
	helper.ResponseCreated(c, "Peminjaman berhasil dibuat", result)
}

func (h *PeminjamanHandler) UpdateStatus(c *gin.Context) {
	var input struct {
		Status string `json:"status" binding:"required"`
	}
	if err := c.ShouldBindJSON(&input); err != nil {
		helper.ResponseValidationError(c, helper.PesanError(err))
		return
	}
	result, err := h.usecase.UpdateStatus(c.Param("id"), input.Status)
	if err != nil {
		helper.ResponseBadRequest(c, err.Error())
		return
	}
	helper.ResponseOK(c, "Status peminjaman berhasil diupdate", result)
}

func (h *PeminjamanHandler) Delete(c *gin.Context) {
	if err := h.usecase.Delete(c.Param("id")); err != nil {
		helper.ResponseNotFound(c, err.Error())
		return
	}
	helper.ResponseOK(c, "Peminjaman berhasil dihapus", nil)
}