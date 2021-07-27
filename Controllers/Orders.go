package Controllers

import (
	"gin-microservice/Config"
	"gin-microservice/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CreateOrder(c *gin.Context) {
	var or Models.Order
	c.BindJSON(&or)
	err := Models.CreateOrder(&or)
	customer :=Models.Customer{CustomerID: or.CustomerCustomerID}
	Config.DB.Create(&customer)

	qty:= or.Quantity
	var pr Models.Product
	er:= Models.GetProductByID(&pr,or.ProductProductID)

	if err == nil && er == nil && qty <= pr.Quantity {
		pr.Quantity -= qty
		Config.DB.Save(pr)
		c.JSON(http.StatusOK, gin.H{
			"order_id":or.OrderID,
			"product_id":or.ProductProductID,
			"qty":or.Quantity,
			"status":"order placed",
		})
	} else {
		c.JSON(http.StatusOK,gin.H{
			"message":"failed",
		})
	}
}
func GetProductsOrderedByCustomer(c *gin.Context) {
	cId := c.Params.ByName("c_id")
	cID,_:=strconv.Atoi(cId)
	var or []Models.Order
	err := Models.GetOrderByCID(&or, int16(cID))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, or)
	}
}
func GetOrderByID(c *gin.Context) {
	oID := c.Params.ByName("o_id")
	OID,_:=strconv.Atoi(oID)
	var or Models.Order
	err := Models.GetOrderByID(&or, int16(OID))
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, or)
	}
}
func GetTraHist(c *gin.Context) {

}



