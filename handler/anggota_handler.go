package handler

import (
    "net/http"
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
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "data": anggotaList})
}

func (h *AnggotaHandler) GetAnggotaByID(c *gin.Context) {
    anggota, err := h.usecase.GetAnggotaByID(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "data": anggota})
}

func (h *AnggotaHandler) SearchAnggota(c *gin.Context) {
    anggotaList, err := h.usecase.SearchAnggota(c.Query("nama_anggota"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "data": anggotaList})
}

func (h *AnggotaHandler) CreateAnggota(c *gin.Context) {
    var anggota model.Anggota
    if err := c.ShouldBindJSON(&anggota); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    result, err := h.usecase.CreateAnggota(anggota)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"success": true, "data": result})
}

func (h *AnggotaHandler) UpdateAnggota(c *gin.Context) {
    var input model.Anggota
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    result, err := h.usecase.UpdateAnggota(c.Param("id"), input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

func (h *AnggotaHandler) DeleteAnggota(c *gin.Context) {
    if err := h.usecase.DeleteAnggota(c.Param("id")); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "message": "Anggota berhasil dihapus"})
}