package models

import "time"

type URL struct {
	ID          uint `gorm:"primaryKey"`
	OriginalURL string
	UID         string `gorm:"index"`
	AddedAt     time.Time
	TTL         time.Time
}
