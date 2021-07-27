package Models

import (
	"gin-microservice/Config"
	_ "github.com/go-sql-driver/mysql"
)
func GetAllProducts(pr *[]Product) (err error) {
	if err = Config.DB.Find(pr).Error; err != nil {
		return err
	}
	return nil
}
func CreateProduct(pr *Product) (err error) {
	err = Config.DB.Create(pr).Error
	if err != nil {
		return err
	}
	return nil
}
func GetProductByID(pr *Product, id int16) (err error) {
	if err = Config.DB.Where("id = ?", id).First(pr).Error; err != nil {
		return err
	}
	return nil
}
func UpdateProduct(pr1 *Product,pr2 *Product) (err error) {
	pr2.Price=pr1.Price
	pr2.Quantity=pr1.Quantity
	Config.DB.Save(pr2)
	return nil
}
func DeleteProduct(pr *Product, id int16) (err error) {
	Config.DB.Where("id = ?", id).Delete(pr)
	return nil
}
