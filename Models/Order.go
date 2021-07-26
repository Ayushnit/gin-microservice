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
func GetStudentByID(or *Order, id string) (err error) {
	if err = Config.DB.Where("id = ?", id).First(or).Error; err != nil {
		return err
	}
	return nil
}

