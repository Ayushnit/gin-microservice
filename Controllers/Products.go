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
		c.JSON(http.StatusOK, pr) //Add message and resolve id issue
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
	var pr1 Models.Product
	id := c.Params.ByName("id")
	err := Models.GetProductByID(&pr1, id)
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
	err := Models.DeleteProduct(&pr, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"id" + id: "is deleted"})
	}
}

