package user

import (
	food2 "OrderingSystem/food"
	order2 "OrderingSystem/order"
	mysql "OrderingSystem/orm"
	"OrderingSystem/shop"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Id int
	Name string
	Password string
	Type int
	Phone string
	Address string
}

func NewUser(name, password, phone, address string) *User{
	return &User{
		Name: name,
		Password: password,
		Phone: phone,
		Address: address,
	}
}

func (u *User)Login() (User, bool){
	db := mysql.NewMysql()
	sql := db.Init()
	var checkUser User
	sql.Find(&checkUser, "name = ?", u.Name)
	return checkUser, u.Password == checkUser.Password
}

func (u *User)Register(){
	db := mysql.NewMysql()
	sql := db.Init()
	sql.AutoMigrate(&User{})
	sql.Create(u)
}

func (u *User)SetShop(s shop.Shop){
	db := mysql.NewMysql()
	sql := db.Init()
	sql.AutoMigrate(&shop.Shop{})
	sql.Create(&s)
}

func (u *User) GetAllShop()[]shop.Shop{
	db := mysql.NewMysql()
	sql := db.Init()
	var shops []shop.Shop
	sql.Find(&shops)
	return shops
}

func (u *User)GetShop(bossId int) shop.Shop{
	db := mysql.NewMysql()
	sql := db.Init()
	var selectShop shop.Shop
	sql.Find(&selectShop, bossId)
	return selectShop
}

func (u *User)AddFood(foodId, shopId int, number int){
	db := mysql.NewMysql()
	sql := db.Init()
	sql.AutoMigrate(&order2.Order{})
	order := order2.Order{UserId: u.Id, FoodId: foodId, ShopId: shopId, Number: number}
	sql.Create(&order)
}

func (u *User)DeleteFood(foodId, shopId int){
	db := mysql.NewMysql()
	sql := db.Init()
	var record order2.Order
	sql.Where("user_id = ? AND food_id = ? AND shop_id= ? AND paid = 0",u.Name, foodId, shopId).First(&record)
	sql.Delete(&record)
}

func (u *User)SelectAllFoods(shopId int) ([]food2.Food, []int) {
	db := mysql.NewMysql()
	sql := db.Init()
	var records []order2.Order
	sql.Where("user_id = ? AND shop_id = ? AND paid = 0", u.Id, shopId).Find(&records)
	foods := make([]food2.Food, len(records))
	numbers :=  make([]int, len(records))
	for index, r := range records {
		var selectFood food2.Food
		sql.First(&selectFood, "id = ?", r.FoodId)
		foods[index] = selectFood
		numbers[index] = r.Number
	}
	return foods, numbers
}

func (u *User)Order(shopId int){
	db := mysql.NewMysql()
	sql := db.Init()
	var records []order2.Order
	sql.Where("user_id = ? AND shop_id AND paid = 0", u.Id, shopId).Find(&records)
	for _, r := range records {
		sql.Model(&r).Update("paid", 1)
	}
}

