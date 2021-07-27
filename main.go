package main

import (
	"fmt"
	"gin-microservice/Config"
	"gin-microservice/Models"
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
	_ = Config.DB.Migrator().DropTable(&Models.Customer{}, &Models.Product{}, &Models.Order{})
	Config.DB.AutoMigrate(&Models.Product{},&Models.Order{},&Models.Customer{})
	//r1 := Routes.SetupRouterRetailer()
	//r1.Run()
	//
	//r2:=Routes.SetupRouterCustomer()
	//r2.Run()
}

