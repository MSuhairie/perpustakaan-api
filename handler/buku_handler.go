package handler

import (
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"perpustakaan-api/helper"
	"perpustakaan-api/model"
	"perpustakaan-api/usecase"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

type BukuHandler struct {
    usecase usecase.BukuUsecase
}

func NewBukuHandler(uc usecase.BukuUsecase) *BukuHandler {
    return &BukuHandler{uc}
}

func (h *BukuHandler) GetAllBuku(c *gin.Context) {
    bukuList, err := h.usecase.GetAllBuku()
    if err != nil {
        helper.ResponseInternalError(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data buku berhasil diambil", bukuList)
}

func (h *BukuHandler) GetBukuByID(c *gin.Context) {
    buku, err := h.usecase.GetBukuByID(c.Param("id"))
    if err != nil {
        helper.ResponseNotFound(c, "Buku tidak ditemukan")
        return
    }
    helper.ResponseOK(c, "Data buku berhasil diambil", buku)
}

func (h *BukuHandler) SearchBuku(c *gin.Context) {
    bukuList, err := h.usecase.SearchBuku(c.Query("judul"))
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Data buku berhasil diambil", bukuList)
}

func (h *BukuHandler) CreateBuku(c *gin.Context) {
    var buku model.Buku
    if err := c.ShouldBindJSON(&buku); err != nil {
        helper.ResponseValidationError(c, helper.PesanError(err))
        return
    }
    result, err := h.usecase.CreateBuku(buku)
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseCreated(c, "Buku berhasil ditambahkan", result)
}

func (h *BukuHandler) UpdateBuku(c *gin.Context) {
    var input model.Buku
    if err := c.ShouldBindJSON(&input); err != nil {
        helper.ResponseValidationError(c, helper.PesanError(err))
        return
    }
    result, err := h.usecase.UpdateBuku(c.Param("id"), input)
    if err != nil {
        helper.ResponseBadRequest(c, err.Error())
        return
    }
    helper.ResponseOK(c, "Buku berhasil diupdate", result)
}

func (h *BukuHandler) DeleteBuku(c *gin.Context) {
    buku, _ := h.usecase.GetBukuByID(c.Param("id"))
    if buku.Foto != "" {
        os.Remove("uploads/" + buku.Foto)
    }
    if err := h.usecase.DeleteBuku(c.Param("id")); err != nil {
        helper.ResponseNotFound(c, "Buku tidak ditemukan")
        return
    }
    helper.ResponseOK(c, "Buku berhasil dihapus", nil)
}

func (h *BukuHandler) UploadFoto(c *gin.Context) {
    // Cari buku dulu
    buku, err := h.usecase.GetBukuByID(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Buku tidak ditemukan"})
        return
    }

    // Ambil file dari request
    file, err := c.FormFile("foto")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "File foto tidak ditemukan"})
        return
    }

    // Validasi ekstensi file
    allowedExt := map[string]bool{
        ".jpg":  true,
        ".jpeg": true,
        ".png":  true,
    }
    ext := strings.ToLower(filepath.Ext(file.Filename))
    if !allowedExt[ext] {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Format file harus jpg, jpeg, atau png",
        })
        return
    }

    // Validasi ukuran file (max 2MB)
    if file.Size > 2*1024*1024 {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Ukuran file maksimal 2MB",
        })
        return
    }

    // Generate nama file unik
    namaFile := fmt.Sprintf("buku_%d_%d%s", buku.ID, time.Now().Unix(), ext)
    pathFile := fmt.Sprintf("uploads/%s", namaFile)

    // Hapus foto lama kalau ada
    if buku.Foto != "" {
        os.Remove("uploads/" + buku.Foto)
    }

    // Simpan file
    if err := c.SaveUploadedFile(file, pathFile); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Gagal menyimpan file",
        })
        return
    }

    // Update nama foto di database
    result, err := h.usecase.UpdateBuku(c.Param("id"), model.Buku{Foto: namaFile})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Foto berhasil diupload",
        "data":    result,
        "url":     fmt.Sprintf("/uploads/%s", namaFile),
    })
}

// Edit/Ganti Foto
func (h *BukuHandler) EditFoto(c *gin.Context) {
    // Cari buku dulu
    buku, err := h.usecase.GetBukuByID(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Buku tidak ditemukan"})
        return
    }

    // Ambil file baru
    file, err := c.FormFile("foto")
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"success": false, "message": "File foto tidak ditemukan"})
        return
    }

    // Validasi ekstensi
    allowedExt := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
    ext := strings.ToLower(filepath.Ext(file.Filename))
    if !allowedExt[ext] {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Format file harus jpg, jpeg, atau png",
        })
        return
    }

    // Validasi ukuran (max 2MB)
    if file.Size > 2*1024*1024 {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Ukuran file maksimal 2MB",
        })
        return
    }

    // Hapus foto lama
    if buku.Foto != "" {
        os.Remove("uploads/" + buku.Foto)
    }

    // Simpan foto baru
    namaFile := fmt.Sprintf("buku_%d_%d%s", buku.ID, time.Now().Unix(), ext)
    pathFile := fmt.Sprintf("uploads/%s", namaFile)
    if err := c.SaveUploadedFile(file, pathFile); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Gagal menyimpan foto",
        })
        return
    }

    // Update nama foto di database
    result, err := h.usecase.UpdateBuku(c.Param("id"), model.Buku{Foto: namaFile})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Foto berhasil diupdate",
        "data":    result,
        "url":     fmt.Sprintf("/uploads/%s", namaFile),
    })
}

// Hapus Foto
func (h *BukuHandler) HapusFoto(c *gin.Context) {
    // Cari buku dulu
    buku, err := h.usecase.GetBukuByID(c.Param("id"))
    if err != nil {
        c.JSON(http.StatusNotFound, gin.H{"success": false, "message": "Buku tidak ditemukan"})
        return
    }

    // Cek apakah ada foto
    if buku.Foto == "" {
        c.JSON(http.StatusBadRequest, gin.H{
            "success": false,
            "message": "Buku tidak memiliki foto",
        })
        return
    }

    // Hapus file dari folder uploads
    if err := os.Remove("uploads/" + buku.Foto); err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{
            "success": false,
            "message": "Gagal menghapus file foto",
        })
        return
    }

    // Kosongkan field foto di database
    result, err := h.usecase.UpdateBuku(c.Param("id"), model.Buku{Foto: ""})
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"success": false, "message": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "success": true,
        "message": "Foto berhasil dihapus",
        "data":    result,
    })
}