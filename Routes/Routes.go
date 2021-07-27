package Routes

import (
	"gin-microservice/Controllers"
	"github.com/gin-gonic/gin"
)
func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/retailer")
	{
		grp1.GET("product", Controllers.GetProducts)
		grp1.POST("product", Controllers.CreateProduct)
		grp1.PATCH("product/:id", Controllers.UpdateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.DELETE("product/:id", Controllers.DeleteProduct)
		grp1.GET("transactionHistory",Controllers.GetTraHist)
		grp1.GET("order/:o_id",Controllers.GetOrderByID)
	}
	grp2:=r.Group("/customer")
	{
		grp2.POST("newCustomer",Controllers.CreateCustomer)
		grp2.GET("product", Controllers.GetProducts)
		grp2.GET("product/:id", Controllers.GetProductByID)
		grp2.POST("order",Controllers.CreateOrder)
		grp2.GET("order/:c_id",Controllers.GetProductsOrderedByCustomer)
	}
	return r
}


