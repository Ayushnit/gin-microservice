package Controllers

import (
"fmt"
"gin-microservice/Models"
"github.com/gin-gonic/gin"
"net/http"
)

func CreateOrder(c *gin.Context) {
	var or Models.Order
	c.BindJSON(&or)
	err := Models.CreateOrder(&or)
	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, or)
	}
}
func GetOrderByID(c *gin.Context) {
	id := c.Params.ByName("id")
	var or Models.Order
	err := Models.GetStudentByID(&or, id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, or)
	}
}



