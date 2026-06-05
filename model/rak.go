package model

import "gorm.io/gorm"

type Rak struct {
    gorm.Model
    KodeRak string `json:"kode_rak" gorm:"uniqueIndex"`
    Lokasi  string `json:"lokasi"`
    Buku    []Buku `json:"buku,omitempty" gorm:"foreignKey:RakID"`
}