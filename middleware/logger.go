package middleware

import (
    "fmt"
    "io"
    "os"
    "time"

    "github.com/gin-gonic/gin"
)

func LoggerMiddleware() gin.HandlerFunc {
    // Buat folder logs kalau belum ada
    os.MkdirAll("logs", os.ModePerm)

    // Buat/buka file log berdasarkan tanggal hari ini
    namaFile := fmt.Sprintf("logs/%s.log", time.Now().Format("2006-01-02"))
    file, err := os.OpenFile(namaFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        fmt.Println("Gagal buat file log:", err)
    }

    // Tulis ke terminal DAN file sekaligus
    multiWriter := io.MultiWriter(os.Stdout, file)

    return gin.LoggerWithConfig(gin.LoggerConfig{
        Formatter: func(param gin.LogFormatterParams) string {
            return fmt.Sprintf("[%s] %s | %d | %s | %s %s\n",
                param.TimeStamp.Format("2006-01-02 15:04:05"),
                param.ClientIP,
                param.StatusCode,
                param.Latency.Round(time.Millisecond),
                param.Method,
                param.Path,
            )
        },
        Output: multiWriter, // ← tulis ke 2 tempat sekaligus
    })
}