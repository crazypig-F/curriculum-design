package food

import (
	"gorm.io/gorm"
)

type Food struct {
	gorm.Model
	Name  string
	Price float32
	MenuId int
}

