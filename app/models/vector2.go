package models

import "github.com/jinzhu/gorm"

type Vector2 struct {
	gorm.Model
	Lat float64
	Lng float64
}
