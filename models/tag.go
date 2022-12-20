package models

type Tag struct {
	ID   uint `gorm:"primaryKey"`
	Name string
}
