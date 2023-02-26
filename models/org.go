package models

type Org struct {
	ID   uint `gorm:"primary_key"`
	Name string
}
