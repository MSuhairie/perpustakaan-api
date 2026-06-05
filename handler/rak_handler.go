package handler

import (
    "net/http"
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
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
        return
    }

    var result []RakResponse
    for _, k := range rakList {
        result = append(result, RakResponse{
            ID:  k.ID,
            KodeRak:  k.KodeRak,
            Lokasi:  k.Lokasi,
        })
    }

    c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

func (h *RakHandler) GetRakByID(c *gin.Context) {
    rak, err := h.usecase.GetRakByID(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
        return
    }

    result := RakResponse{
        ID:           rak.ID,
        KodeRak: rak.KodeRak,
        Lokasi: rak.Lokasi,
    }

    c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

func (h *RakHandler) SearchRak(c *gin.Context) {
    rakList, err := h.usecase.SearchRak(c.Query("nama_rak"))
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }

    var result []RakResponse
    for _, k := range rakList {
        result = append(result, RakResponse{
            ID:  k.ID,
            KodeRak:  k.KodeRak,
            Lokasi:  k.Lokasi,
        })
    }

    c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

func (h *RakHandler) CreateRak(c *gin.Context) {
    var rak model.Rak
    if err := c.ShouldBindJSON(&rak); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    result, err := h.usecase.CreateRak(rak)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"success": true, "data": result})
}

func (h *RakHandler) UpdateRak(c *gin.Context) {
    var input model.Rak
    if err := c.ShouldBindJSON(&input); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    result, err := h.usecase.UpdateRak(c.Param("id"), input)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "data": result})
}

func (h *RakHandler) DeleteRak(c *gin.Context) {
    if err := h.usecase.DeleteRak(c.Param("id")); err != nil {
        c.JSON(http.StatusNotFound, gin.H{"success": false, "message": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true, "message": "Rak berhasil dihapus"})
}