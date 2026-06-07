package handler

import (
    "perpustakaan-api/helper"
    "perpustakaan-api/usecase"

    "github.com/gin-gonic/gin"
)

type PengembalianHandler struct {
    usecase usecase.PengembalianUsecase
}

func NewPengembalianHandler(uc usecase.PengembalianUsecase) *PengembalianHandler {
    return &PengembalianHandler{uc}
}

func (h *PengembalianHandler) GetAll(c *gin.Context) {
    list, err := h.usecase.GetAll()
    if err != nil {
        helper.ResponseInternalError(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data pengembalian berhasil diambil", list)
}

func (h *PengembalianHandler) GetByID(c *gin.Context) {
    p, err := h.usecase.GetByID(c.Param("id"))
    if err != nil {
        helper.ResponseNotFound(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data pengembalian berhasil diambil", p)
}

func (h *PengembalianHandler) Create(c *gin.Context) {
    var input usecase.CreatePengembalianInput
    if err := c.ShouldBindJSON(&input); err != nil {
        helper.ResponseValidationError(c, helper.PesanError(err))
        return
    }
    result, err := h.usecase.Create(input)
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseCreated(c, "Pengembalian berhasil diproses", result)
}