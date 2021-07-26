package main

import (
	"fmt"
	"gin-microservice/Config"
	"gin-microservice/Models"
	"gin-microservice/Routes"
	"github.com/jinzhu/gorm"
)
var err error
func main() {
	Config.DB, err = gorm.Open("mysql", Config.DbURL(Config.BuildDBConfig()))
	if err != nil {
		fmt.Println("Status:", err)
	}
	defer Config.DB.Close()
	Config.DB.AutoMigrate(&Models.Product{})
	r1 := Routes.SetupRouterProduct()
	//running
	r1.Run()

	r2:=Routes.SetupRouterOrder()
	r2.Run()
}

