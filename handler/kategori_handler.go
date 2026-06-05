package handler

import (
    "net/http"
    "perpustakaan-api/model"
    "perpustakaan-api/usecase"

    "github.com/gin-gonic/gin"
)

type KategoriResponse struct {
    ID uint `json:"id"`
    NamaKategori  string  `json:"nama_kategori"`
}

type KategoriHandler struct {
    usecase usecase.KategoriUsecase
}

func NewKategoriHandler(uc usecase.KategoriUsecase) *KategoriHandler {
    return &KategoriHandler{uc}
}

func (h *KategoriHandler) GetAllKategori(c *gin.Context) {
    kategoriList, err := h.usecase.GetAllKategori()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
        return
    }

    var result []KategoriResponse
    for _, k := range kategoriList {
        result = append(result, KategoriResponse{
            ID:  k.ID,
            NamaKategori:  k.NamaKategori,
        })
    }

    c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

func (h *KategoriHandler) GetKategoriByID(c *gin.Context) {
    kategori, err := h.usecase.GetKategoriByID(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
        return
    }

    result := KategoriResponse{
        ID:           kategori.ID,
        NamaKategori: kategori.NamaKategori,
    }

    c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

func (h *KategoriHandler) SearchKategori(c *gin.Context) {
    kategoriList, err := h.usecase.SearchKategori(c.Query("nama_kategori"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }

    var result []KategoriResponse
    for _, k := range kategoriList {
        result = append(result, KategoriResponse{
            ID:  k.ID,
            NamaKategori:  k.NamaKategori,
        })
    }

    c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

func (h *KategoriHandler) CreateKategori(c *gin.Context) {
    var kategori model.Kategori
    if err := c.ShouldBindJSON(&kategori); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    result, err := h.usecase.CreateKategori(kategori)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"success": true, "data": result})
}

func (h *KategoriHandler) UpdateKategori(c *gin.Context) {
    var input model.Kategori
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    result, err := h.usecase.UpdateKategori(c.Param("id"), input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

func (h *KategoriHandler) DeleteKategori(c *gin.Context) {
    if err := h.usecase.DeleteKategori(c.Param("id")); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "message": "Kategori berhasil dihapus"})
}