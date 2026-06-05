package helper

import (
    "errors"
    "fmt"
    "strings"

    "github.com/go-playground/validator/v10"
    "gorm.io/gorm"
)

func PesanError(err error) map[string]string {
    errorsMap := make(map[string]string)

    var validationErrors validator.ValidationErrors
    if errors.As(err, &validationErrors) {
        for _, fieldError := range validationErrors {
            field := fieldError.Field()
            switch fieldError.Tag() {
            case "required":
                errorsMap[field] = fmt.Sprintf("%s wajib diisi", field)
            case "email":
                errorsMap[field] = "Format email tidak valid"
            case "min":
                errorsMap[field] = fmt.Sprintf("%s minimal %s karakter", field, fieldError.Param())
            case "max":
                errorsMap[field] = fmt.Sprintf("%s maksimal %s karakter", field, fieldError.Param())
            case "numeric":
                errorsMap[field] = fmt.Sprintf("%s harus berupa angka", field)
            case "gte":
                errorsMap[field] = fmt.Sprintf("%s harus lebih dari atau sama dengan %s", field, fieldError.Param())
            case "lte":
                errorsMap[field] = fmt.Sprintf("%s harus kurang dari atau sama dengan %s", field, fieldError.Param())
            case "gt":
                errorsMap[field] = fmt.Sprintf("%s harus lebih dari %s", field, fieldError.Param())
            case "lt":
                errorsMap[field] = fmt.Sprintf("%s harus kurang dari %s", field, fieldError.Param())
            case "url":
                errorsMap[field] = fmt.Sprintf("%s harus berupa URL yang valid", field)
            case "uuid":
                errorsMap[field] = fmt.Sprintf("%s harus berupa UUID yang valid", field)
            default:
                errorsMap[field] = "Nilai tidak valid"
            }
        }
    }

    if err != nil {
        if JikaDuplikat(err) {
            switch {
            case strings.Contains(err.Error(), "username"):
                errorsMap["username"] = "Username sudah digunakan"
            case strings.Contains(err.Error(), "email"):
                errorsMap["email"] = "Email sudah digunakan"
            case strings.Contains(err.Error(), "no_hp"):
                errorsMap["no_hp"] = "Nomor HP sudah digunakan"
            case strings.Contains(err.Error(), "isbn"):
                errorsMap["isbn"] = "ISBN sudah digunakan"
            default:
                errorsMap["duplicate"] = "Data sudah ada"
            }
        } else if errors.Is(err, gorm.ErrRecordNotFound) {
            errorsMap["error"] = "Data tidak ditemukan"
        }
    }

    return errorsMap
}

// PostgreSQL pakai "duplicate key value violates unique constraint"
func JikaDuplikat(err error) bool {
    return err != nil && strings.Contains(err.Error(), "duplicate key value violates unique constraint")
}