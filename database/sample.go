package database

import "github.com/arut-ji/individual-project/sample"

type Sample struct {
	ID     uint          `gorm:"primaryKey"`
	Sample sample.Sample `gorm:"embedded"`
}
