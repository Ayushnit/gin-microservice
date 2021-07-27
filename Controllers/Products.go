package Controllers

import (
	"fmt"
	"gin-microservice/Models"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetProducts(c *gin.Context) {
	var pr []Models.Product
	err := Models.GetAllProducts(&pr)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pr)
	}
}
func CreateProduct(c *gin.Context) {
	var pr Models.Product
	c.BindJSON(&pr)
	err := Models.CreateProduct(&pr)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{
			"product_id":pr.ProductID,
			"name":pr.Name,
			"price":pr.Price,
			"qty":pr.Quantity,
			"message":"product successfully added",
		})
	}
}
func GetProductByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var pr Models.Product
	err := Models.GetProductByID(&pr, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pr)
	}
}
func UpdateProduct(c *gin.Context) {
	var pr2 Models.Product
	c.BindJSON(&pr2)
	qty:=pr2.Quantity
	price:=pr2.Price

	id := c.Params.ByName("id")
	err := Models.GetProductByID(&pr2,id)
	er:=Models.UpdateProduct(&pr2,qty,price)
	if err != nil || er !=nil{
		c.JSON(http.StatusNotFound, pr2)
	} else {
		c.JSON(http.StatusOK,pr2)
	}
}
func DeleteProduct(c *gin.Context) {
	var pr Models.Product
	id := c.Params.ByName("id")
	err := Models.DeleteProduct(&pr,id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

