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
func GetProductByID(pr *Product, id string) (err error) {
	if err = Config.DB.Where("product_id = ?", id).First(pr).Error; err != nil {
		return err
	}
	return nil
}
func UpdateProduct(pr2 *Product,qty int16,price int16) (err error) {
	pr2.Price=price
	pr2.Quantity=qty
	Config.DB.Save(pr2)
	return nil
}
func DeleteProduct(pr *Product, id string) (err error) {
	Config.DB.Where("product_id = ?", id).Delete(pr)
	return nil
}
