package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type Deformation3 struct {
	gorm.Model
	Position Vector3
	Deformation float32
	Time time.Time
}
