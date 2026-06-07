package handler

import (
	"perpustakaan-api/dto"
	"perpustakaan-api/helper"
	"perpustakaan-api/model"
	"perpustakaan-api/usecase"

	"github.com/gin-gonic/gin"
)

type RakResponse struct {
    ID uint `json:"id"`
    KodeRak  string  `json:"kode_rak"`
    Lokasi  string  `json:"lokasi"`
}

type RakHandler struct {
    usecase usecase.RakUsecase
}

func NewRakHandler(uc usecase.RakUsecase) *RakHandler {
    return &RakHandler{uc}
}

func (h *RakHandler) GetAllRak(c *gin.Context) {
    rakList, err := h.usecase.GetAllRak()
    if err != nil {
        helper.ResponseInternalError(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data Rak berhasil diambil", helper.ToRakResponseList(rakList))
}

func (h *RakHandler) GetRakByID(c *gin.Context) {
    rak, err := h.usecase.GetRakByID(c.Param("id"))
    if err != nil {
        helper.ResponseNotFound(c, "Rak tidak ditemukan")
        return
    }
    helper.ResponseOK(c, "Data Rak berhasil diambil", helper.ToRakResponse(rak))
}

func (h *RakHandler) SearchRak(c *gin.Context) {
    rakList, err := h.usecase.SearchRak(c.Query("nama_rak"))
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data Rak berhasil diambil", helper.ToRakResponseList(rakList))
}

func (h *RakHandler) CreateRak(c *gin.Context) {
    var req dto.RakRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseValidationError(c, helper.PesanError(err))
        return
    }

    rak := model.Rak{
        KodeRak: req.KodeRak,
        Lokasi: req.Lokasi,
    }

    result, err := h.usecase.CreateRak(rak)
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
	helper.ResponseCreated(c, "Rak berhasil ditambahkan", helper.ToRakResponse(result))
}

func (h *RakHandler) UpdateRak(c *gin.Context) {
    var req dto.RakRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseValidationError(c, helper.PesanError(err))
        return
    }

    input := model.Rak{
        KodeRak: req.KodeRak,
        Lokasi: req.Lokasi,
    }

    result, err := h.usecase.UpdateRak(c.Param("id"), input)
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data Rak berhasil diupdate", helper.ToRakResponse(result))

}

func (h *RakHandler) DeleteRak(c *gin.Context) {
    if err := h.usecase.DeleteRak(c.Param("id")); err != nil {
        helper.ResponseNotFound(c, "Rak tidak ditemukan")
        return
    }
    helper.ResponseOK(c, "Rak berhasil dihapus", nil)
}