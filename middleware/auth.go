package middleware

import (
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func GetSecretKey() []byte {
    secret := os.Getenv("JWT_SECRET")
    if secret == "" {
        secret = "secret-perpustakaan" // fallback
    }
    return []byte(secret)
}

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        // Ambil token dari header
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.JSON(http.StatusUnauthorized, gin.H{
                "success": false,
                "message": "Token tidak ditemukan",
            })
            c.Abort()
            return
        }

        // Format: "Bearer <token>" → ambil tokennya saja
        tokenString := strings.Replace(authHeader, "Bearer ", "", 1)

        // Validasi token
        token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
            return GetSecretKey(), nil
        })

        if err != nil || !token.Valid {
            c.JSON(http.StatusUnauthorized, gin.H{
                "success": false,
                "message": "Token tidak valid",
            })
            c.Abort()
            return
        }

        // Simpan data user ke context
        claims := token.Claims.(jwt.MapClaims)
        c.Set("user_id", claims["user_id"])
        c.Set("email", claims["email"])

        c.Next() // Lanjut ke handler
    }
}

