package shop

import (
	food2 "OrderingSystem/food"
	"OrderingSystem/menu"
	mysql "OrderingSystem/orm"
	"gorm.io/gorm"
)

type Shop struct {
	gorm.Model
	Id int
	Name string
	BossId int
	MenuName string
}

func (s *Shop) GetMenu() menu.Menu{
	db := mysql.NewMysql()
	sql := db.Init()
	var SelectMenu menu.Menu
	sql.Find(&SelectMenu, "shop_id = ?", s.Id)
	return SelectMenu
}

func (s *Shop) SetMenu(m menu.Menu){
	db := mysql.NewMysql()
	sql := db.Init()
	sql.AutoMigrate(&menu.Menu{})
	sql.Create(&m)
}

func (s *Shop) SaveFood(food food2.Food) {
	menuId := s.GetMenu().Id
	food.MenuId = menuId
	db := mysql.NewMysql()
	sql := db.Init()
	sql.AutoMigrate(&food2.Food{})
	sql.Create(&food)
}

func (s *Shop) DeleteFood(foodName string) {
	db := mysql.NewMysql()
	sql := db.Init()
	var food food2.Food
	sql.First(&food, "name = ?", foodName)
	sql.Delete(&food)
}

func (s *Shop) UpdateFood(foodName string, price float32){
	db := mysql.NewMysql()
	sql := db.Init()
	var food food2.Food
	sql.Find(&food, "name = ?", foodName)
	sql.Model(&food).Update("price", price)
}

func (s *Shop) SelectAllFoods() []food2.Food{
	db := mysql.NewMysql()
	sql := db.Init()
	var foods []food2.Food
	sql.Find(&foods)
	return foods
}

func (s *Shop) GetFood(foodId int) food2.Food{
	db := mysql.NewMysql()
	sql := db.Init()
	var food food2.Food
	sql.First(&food, "id = ?", foodId)
	return food
}


