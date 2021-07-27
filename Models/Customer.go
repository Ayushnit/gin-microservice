package Models

import (
	"gin-microservice/Config"
_ "github.com/go-sql-driver/mysql"
)

func CreateCustomer1(cu *Customer) (err error) {
	if err = Config.DB.Create(cu).Error; err != nil {
		return err
	}
	return nil
}
