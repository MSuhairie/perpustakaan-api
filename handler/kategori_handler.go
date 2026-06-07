package handler

import (
	"perpustakaan-api/dto"
	"perpustakaan-api/helper"
	"perpustakaan-api/model"
	"perpustakaan-api/usecase"

	"github.com/gin-gonic/gin"
)

type KategoriHandler struct {
    usecase usecase.KategoriUsecase
}

func NewKategoriHandler(uc usecase.KategoriUsecase) *KategoriHandler {
    return &KategoriHandler{uc}
}

func (h *KategoriHandler) GetAllKategori(c *gin.Context) {
    kategoriList, err := h.usecase.GetAllKategori()
    if err != nil {
        helper.ResponseInternalError(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data Kategori berhasil diambil", helper.ToKategoriResponseList(kategoriList))
}

func (h *KategoriHandler) GetKategoriByID(c *gin.Context) {
    kategori, err := h.usecase.GetKategoriByID(c.Param("id"))
    if err != nil {
        helper.ResponseNotFound(c, "Kategori tidak ditemukan")
        return
    }
    helper.ResponseOK(c, "Data Kategori berhasil diambil", helper.ToKategoriResponse(kategori))
}

func (h *KategoriHandler) SearchKategori(c *gin.Context) {
    kategoriList, err := h.usecase.SearchKategori(c.Query("nama_kategori"))
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data Kategori berhasil diambil", helper.ToKategoriResponseList(kategoriList))
}

func (h *KategoriHandler) CreateKategori(c *gin.Context) {
    var req dto.KategoriRequest
    if err := c.ShouldBindJSON(&req); err != nil {
		helper.ResponseValidationError(c, helper.PesanError(err))
        return
    }
    kategori := model.Kategori{
        NamaKategori: req.NamaKategori,
    }
    result, err := h.usecase.CreateKategori(kategori)
    if err != nil {
		helper.ResponseBadRequest(c, err.Error())
        return
    }
	helper.ResponseCreated(c, "Kategori berhasil ditambahkan", helper.ToKategoriResponse(result))
}

func (h *KategoriHandler) UpdateKategori(c *gin.Context) {
    var req dto.KategoriRequest
    if err := c.ShouldBindJSON(&req); err != nil {
        helper.ResponseValidationError(c, helper.PesanError(err))
        return
    }
    input := model.Kategori{
        NamaKategori: req.NamaKategori,
    }
    result, err := h.usecase.UpdateKategori(c.Param("id"), input)
    if err != nil {
		helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data Kategori berhasil diupdate", helper.ToKategoriResponse(result))
}

func (h *KategoriHandler) DeleteKategori(c *gin.Context) {
    if err := h.usecase.DeleteKategori(c.Param("id")); err != nil {
        helper.ResponseNotFound(c, "Kategori tidak ditemukan")
        return
    }
    helper.ResponseOK(c, "Kategori berhasil dihapus", nil)

}