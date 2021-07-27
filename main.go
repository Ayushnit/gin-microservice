package main

import (
	"fmt"
	"gin-microservice/Config"
	"gin-microservice/Models"
	"gin-microservice/Routes"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)
var err error
func main() {
	Config.DB, err = gorm.Open(mysql.Open(Config.DbURL(Config.BuildDBConfig())), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		fmt.Println("Status:", err)
	}
	Config.DB.AutoMigrate(&Models.Product{},&Models.Order{},&Models.Customer{})

	r := Routes.SetupRouter()
	r.Run()

}

