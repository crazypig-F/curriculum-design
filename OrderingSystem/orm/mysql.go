package mysql

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	SqlType  string
	UserName string
	Password string
	DBName   string
}

func NewMysql() *Mysql {
	return &Mysql{
		SqlType:  "mysql",
		UserName: "root",
		Password: "123",
		DBName:   "dinner",
	}
}

func (m *Mysql) Init() *gorm.DB {
	db, err := gorm.Open(mysql.Open(m.UserName+":"+m.Password+"@(localhost)/"+m.DBName+"?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
	}
	return db
}



