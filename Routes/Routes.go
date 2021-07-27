package Routes

import (
	"gin-microservice/Controllers"
	"github.com/gin-gonic/gin"
)
func SetupRouterRetailer() *gin.Engine {
	r1 := gin.Default()
	grp1 := r1.Group("/retailer")
	{
		grp1.GET("product", Controllers.GetProducts)
		grp1.POST("product", Controllers.CreateProduct)
		grp1.PATCH("product/:id", Controllers.UpdateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.DELETE("product/:id", Controllers.DeleteProduct)
		grp1.GET("transactionHistory",Controllers.GetTraHist)
		grp1.GET("order/:o_id",Controllers.GetOrderByID)
	}

	return r1
}
func SetupRouterCustomer() *gin.Engine {
	r2:=gin.Default()
	grp2:=r2.Group("/customer")
	{
		grp2.GET("product", Controllers.GetProducts)
		grp2.GET("product/{id}", Controllers.GetProductByID)
		grp2.POST("order",Controllers.CreateOrder)
		grp2.GET("order/:c_id",Controllers.GetProductsOrderedByCustomer)
	}
	return r2
}


