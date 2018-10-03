package models

import "github.com/jinzhu/gorm"

type Vector3 struct {
	gorm.Model
	X float64
	Y float64
	Z float64
}
