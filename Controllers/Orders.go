package Controllers

import (
	"fmt"
	"gin-microservice/Config"
	"gin-microservice/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func CreateOrder(c *gin.Context) {
	var ord Models.Order
	c.BindJSON(&ord)

	var cu Models.Customer

	er1:=Config.DB.Where("customer_id = ?",ord.CustomerCustomerID).First(&cu).Error
	if er1 !=nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	status:=cu.Active
	fmt.Println(status)
	fmt.Println(ord.Quantity)
	fmt.Println(ord.CustomerCustomerID)
	fmt.Println(ord.ProductProductID)

	var flag bool=true
	qty:= ord.Quantity
	var pr Models.Product
	er:= Models.GetProductByID(&pr, strconv.FormatUint(uint64(ord.ProductProductID),10))
	if er !=nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	if er==nil && er1==nil && qty <= pr.Quantity && status==true{
		err := Models.CreateOrder(&ord)
		if err==nil {
			pr.Quantity -= qty
			Config.DB.Save(pr)
			Config.DB.Model(&cu).Where("customer_id = ?", ord.CustomerCustomerID).Update("active", false)
			time.Sleep(200 * time.Second)
			Config.DB.Model(&cu).Where("customer_id = ?", ord.CustomerCustomerID).Update("active", true)
		} else {
			flag=false
			c.AbortWithStatus(http.StatusNotFound)
		}
	} else {
		flag=false
		c.AbortWithStatus(http.StatusNotFound)
	}
	if flag==true {
		c.JSON(http.StatusOK, gin.H{
			"order_id":   ord.OrderID,
			"product_id": ord.ProductProductID,
			"qty":        ord.Quantity,
			"status":     "order placed",
		})
	} else {
		c.JSON(http.StatusNotFound,gin.H{
			"status":"failed",
		})
	}
}

func GetProductsOrderedByCustomer(c *gin.Context) {
	cId := c.Params.ByName("c_id")
	var or []Models.Order
	err := Models.GetOrderByCID(&or,cId)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, or)
	}
}
func GetOrderByID(c *gin.Context) {
	oID := c.Params.ByName("o_id")
	var or Models.Order
	err := Models.GetOrderByID(&or, oID)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, or)
	}
}
func GetTraHist(c *gin.Context) {

}
func CreateCustomer(c *gin.Context) {
	var cu Models.Customer
	c.BindJSON(&cu)

	err:=Models.CreateCustomer1(&cu)

	if err != nil {
		fmt.Println(err.Error())
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK,cu)
	}
}


