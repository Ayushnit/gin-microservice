package Models

import (
	"gin-microservice/Config"
	_ "github.com/go-sql-driver/mysql"
)

func CreateOrder(or *Order) (err error) {
	if err = Config.DB.Create(or).Error; err != nil {
		return err
	}
	return nil
}
func GetOrderByID(or *Order, oID int16) (err error) {
	if err = Config.DB.Where("order_id = ?", oID).First(or).Error; err != nil {
		return err
	}
	return nil
}
func GetOrderByCID(or *[]Order, cID int16) (err error) {
	if err = Config.DB.Where("customer_id = ?", cID).Find(or).Error; err != nil {
		return err
	}
	return nil
}

