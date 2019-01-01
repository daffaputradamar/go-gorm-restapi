package config

import (
	"os"

	"github.com/daffaputradamar/go-gorm-restapi/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func ConnectDB() *gorm.DB {
	db, err := gorm.Open("mysql", os.Getenv("DB_USER")+":"+os.Getenv("DB_PASS")+"@/"+os.Getenv("DB_NAME")+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("Failed to connect")
	}

	db.AutoMigrate(&models.Post{})
	return db
}
