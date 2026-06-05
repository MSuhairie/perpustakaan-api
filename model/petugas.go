package model

import "gorm.io/gorm"

type Petugas struct {
    gorm.Model
    Nama     string `json:"nama"`
    Username string `json:"username" gorm:"unique"`
    Password string `json:"-"`
    Role     string `json:"role"`
}