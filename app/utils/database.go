package utils

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sigit-khoirun-nizam/go-fiber-onlinestore/app/models"
)

var DB *gorm.DB

func InitDatabase() {
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/synapsis?parseTime=true")
	if err != nil {
		panic("Gagal Menghubungkan Ke Database")
	}
	DB = db
	DB.AutoMigrate(&models.Product{}, &models.CartItem{}, &models.User{})
}
