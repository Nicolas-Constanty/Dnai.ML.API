package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Deformation2 struct {
	gorm.Model
	Position Vector2
	Deformation float32
	Time time.Time
}
