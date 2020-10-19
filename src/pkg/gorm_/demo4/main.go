package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 关联一对多
// https://gorm.io/zh_CN/docs/has_many.html

var db *gorm.DB

func init() {
	var err error
	dns := "root:123456@tcp(127.0.0.1:3306)/blog?charset=utf8&parseTime=True&loc=Local&timeout=10s"
	db, err = gorm.Open(mysql.Open(dns), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

type User struct {
	gorm.Model
	Name      string
	Languages []Language `gorm:"many2many:user_languages;"`
}

type Language struct {
	gorm.Model
	Name string
}

func run() {
	// ===============创建数据=========
	db.AutoMigrate(&User{})
	user := User{
		Name: "jinzhu",
		Languages: []Language{
			{Name: "ZH"},
			{Name: "EN"},
		},
	}

	db.Create(&user)
	//db.Create(u1)
	// ===============查询数据=========
}
func main() {
	run()
}
