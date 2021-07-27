package Controllers

import (
	"fmt"
	"gin-microservice/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
	ID,_:=strconv.Atoi(id)
	var pr Models.Product
	err := Models.GetProductByID(&pr, int16(ID))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pr)
	}
}
func UpdateProduct(c *gin.Context) {
	var pr1 Models.Product
	id := c.Params.ByName("id")
	ID,_:=strconv.Atoi(id)
	err := Models.GetProductByID(&pr1, int16(ID))
	if err != nil {
		c.JSON(http.StatusNotFound, pr1)
	}
	c.BindJSON(&pr1)
	var pr2 Models.Product
	err = Models.UpdateProduct(&pr1,&pr2)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, pr1)
	}
}
func DeleteProduct(c *gin.Context) {
	var pr Models.Product
	id := c.Params.ByName("id")
	ID,_:=strconv.Atoi(id)
	err := Models.DeleteProduct(&pr, int16(ID))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

