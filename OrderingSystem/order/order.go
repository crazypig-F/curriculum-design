package order

import (
	"gorm.io/gorm"
)

type Order struct {
	gorm.Model
	Id int
	UserId int
	FoodId int
	ShopId int
	Number int
	Paid bool
}
