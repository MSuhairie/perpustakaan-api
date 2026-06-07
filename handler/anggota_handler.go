package handler

import (
	"perpustakaan-api/dto"
	"perpustakaan-api/helper"
	"perpustakaan-api/model"
	"perpustakaan-api/usecase"

	"github.com/gin-gonic/gin"
)

type AnggotaHandler struct {
    usecase usecase.AnggotaUsecase
}

func NewAnggotaHandler(uc usecase.AnggotaUsecase) *AnggotaHandler {
    return &AnggotaHandler{uc}
}

func (h *AnggotaHandler) GetAllAnggota(c *gin.Context) {
    anggotaList, err := h.usecase.GetAllAnggota()
    if err != nil {
        helper.ResponseInternalError(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data anggota berhasil diambil", helper.ToAnggotaResponseList(anggotaList))
}

func (h *AnggotaHandler) GetAnggotaByID(c *gin.Context) {
    anggota, err := h.usecase.GetAnggotaByID(c.Param("id"))
    if err != nil {
        helper.ResponseNotFound(c, "Buku tidak ditemukan")
        return
    }
    helper.ResponseOK(c, "Data anggota berhasil diambil", helper.ToAnggotaResponse(anggota))
}

func (h *AnggotaHandler) SearchAnggota(c *gin.Context) {
    anggotaList, err := h.usecase.SearchAnggota(c.Query("nama_anggota"))
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data anggota berhasil diambil", helper.ToAnggotaResponseList(anggotaList))
}

func (h *AnggotaHandler) CreateAnggota(c *gin.Context) {
    var req dto.AnggotaRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseValidationError(c, helper.PesanError(err))
        return
    }
    anggota := model.Anggota{
        Nama: req.Nama,
        Alamat: req.Alamat,
        NoHP: req.NoHP,
        TglDaftar: req.TglDaftar,
        Status: req.Status,
    }
    result, err := h.usecase.CreateAnggota(anggota)
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseCreated(c, "Anggota berhasil ditambahkan", helper.ToAnggotaResponse(result))
}

func (h *AnggotaHandler) UpdateAnggota(c *gin.Context) {
    var req dto.AnggotaRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseValidationError(c, helper.PesanError(err))
        return
    }
    input := model.Anggota{
        Nama: req.Nama,
        Alamat: req.Alamat,
        NoHP: req.NoHP,
        TglDaftar: req.TglDaftar,
        Status: req.Status,
    }
    result, err := h.usecase.UpdateAnggota(c.Param("id"), input)
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Anggota berhasil diupdate", helper.ToAnggotaResponse(result))
}

func (h *AnggotaHandler) DeleteAnggota(c *gin.Context) {
    if err := h.usecase.DeleteAnggota(c.Param("id")); err != nil {
        helper.ResponseNotFound(c, "Anggota tidak ditemukan")
        return
    }
    helper.ResponseOK(c, "Anggota berhasil dihapus", nil)
}