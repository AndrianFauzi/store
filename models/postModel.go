package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	ID uint
	COdeProduct string
	NamaProduct string
	HargaProduct int64
	UOM string
	CreatedAt time.Time
	UpdatedAt time.Time
}