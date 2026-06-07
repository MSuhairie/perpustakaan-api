package main

import (
	"os"
	"perpustakaan-api/database"
	"perpustakaan-api/handler"
	"perpustakaan-api/middleware"
	"perpustakaan-api/model"
	"perpustakaan-api/repository"
	"perpustakaan-api/usecase"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    godotenv.Load()
    database.Connect()

    database.DB.AutoMigrate(
        &model.Kategori{},
        &model.Rak{},
        &model.Buku{},
        &model.Anggota{},
        &model.Petugas{},
        &model.Peminjaman{},
        &model.DetailPeminjaman{},
        &model.Pengembalian{},
        &model.DetailPengembalian{},
        &model.Denda{},
    )

    // Dependency Injection
    bukuRepo    := repository.NewBukuRepository(database.DB)
    bukuUsecase := usecase.NewBukuUsecase(bukuRepo)
    bukuHandler := handler.NewBukuHandler(bukuUsecase)

    kategoriRepo    := repository.NewKategoriRepository(database.DB)
    kategoriUsecase := usecase.NewKategoriUsecase(kategoriRepo)
    kategoriHandler := handler.NewKategoriHandler(kategoriUsecase)

    rakRepo    := repository.NewRakRepository(database.DB)
    rakUsecase := usecase.NewRakUsecase(rakRepo)
    rakHandler := handler.NewRakHandler(rakUsecase)

    anggotaRepo    := repository.NewAnggotaRepository(database.DB)
    anggotaUsecase := usecase.NewAnggotaUsecase(anggotaRepo)
    anggotaHandler := handler.NewAnggotaHandler(anggotaUsecase)

    peminjamanRepo    := repository.NewPeminjamanRepository(database.DB)
    peminjamanUsecase := usecase.NewPeminjamanUsecase(peminjamanRepo, bukuRepo)
    peminjamanHandler := handler.NewPeminjamanHandler(peminjamanUsecase)

    dendaRepo           := repository.NewDendaRepository(database.DB)
    pengembalianRepo    := repository.NewPengembalianRepository(database.DB)
    pengembalianUsecase := usecase.NewPengembalianUsecase(pengembalianRepo, peminjamanRepo, bukuRepo, dendaRepo)
    pengembalianHandler := handler.NewPengembalianHandler(pengembalianUsecase)

    r := gin.New()

    r.Use(middleware.LoggerMiddleware()) // ← Logger
    r.Use(gin.Recovery())                // ← Recovery (tangkap panic)
    r.Use(middleware.CorsMiddleware())   // ← CORS

    // Static files — akses foto lewat URL
    r.Static("/uploads", "./uploads")

    r.GET("/", func(c *gin.Context) {
        c.JSON(200, gin.H{"message": "Perpustakaan API", "version": "1.0.0", "status":  "running"})
    })

    r.POST("/register", handler.Register)
    r.POST("/login", handler.Login)

    auth := r.Group("/", middleware.AuthMiddleware())
    {
        // Buku routes
        buku := auth.Group("/buku")
        {
            buku.GET("", bukuHandler.GetAllBuku)
            buku.GET("/search", bukuHandler.SearchBuku)
            buku.GET("/:id", bukuHandler.GetBukuByID)
            buku.POST("", bukuHandler.CreateBuku)
            buku.PUT("/:id", bukuHandler.UpdateBuku)
            buku.DELETE("/:id", bukuHandler.DeleteBuku)
            buku.POST("/:id/foto", bukuHandler.UploadFoto)
            buku.PUT("/:id/foto", bukuHandler.EditFoto)
            buku.DELETE("/:id/foto", bukuHandler.HapusFoto)
        }
        
        kategori := auth.Group("/kategori")
        {
            kategori.GET("", kategoriHandler.GetAllKategori)
            kategori.GET("/search", kategoriHandler.SearchKategori)
            kategori.GET("/:id", kategoriHandler.GetKategoriByID)
            kategori.POST("", kategoriHandler.CreateKategori)
            kategori.PUT("/:id", kategoriHandler.UpdateKategori)
            kategori.DELETE("/:id", kategoriHandler.DeleteKategori)
        }

        rak := auth.Group("/rak")
        {
            rak.GET("", rakHandler.GetAllRak)
            rak.GET("/search", rakHandler.SearchRak)
            rak.GET("/:id", rakHandler.GetRakByID)
            rak.POST("", rakHandler.CreateRak)
            rak.PUT("/:id", rakHandler.UpdateRak)
            rak.DELETE("/:id", rakHandler.DeleteRak)
        }

        anggota := auth.Group("/anggota")
        {
            anggota.GET("", anggotaHandler.GetAllAnggota)
            anggota.GET("/search", anggotaHandler.SearchAnggota)
            anggota.GET("/:id", anggotaHandler.GetAnggotaByID)
            anggota.POST("", anggotaHandler.CreateAnggota)
            anggota.PUT("/:id", anggotaHandler.UpdateAnggota)
            anggota.DELETE("/:id", anggotaHandler.DeleteAnggota)
        }

        peminjaman := auth.Group("/peminjaman")
        {
            peminjaman.GET("", peminjamanHandler.GetAll)
            peminjaman.GET("/:id", peminjamanHandler.GetByID)
            peminjaman.POST("", peminjamanHandler.Create)
            peminjaman.PUT("/:id/status", peminjamanHandler.UpdateStatus)
            peminjaman.DELETE("/:id", peminjamanHandler.Delete)
        }

        pengembalian := auth.Group("/pengembalian")
        {
            pengembalian.GET("", pengembalianHandler.GetAll)
            pengembalian.GET("/:id", pengembalianHandler.GetByID)
            pengembalian.POST("", pengembalianHandler.Create)
        }

    }

    port := os.Getenv("APP_PORT")
    if port == "" {
        port = "8080"
    }
    r.Run(":" + port)
}