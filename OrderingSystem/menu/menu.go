package menu

import "gorm.io/gorm"

type Menu struct {
	gorm.Model
	Id int
	Name string
	ShopId int
}