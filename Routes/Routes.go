package Routes

import (
	"gin-microservice/Controllers"
	"github.com/gin-gonic/gin"
)
func SetupRouterProduct() *gin.Engine {
	r1 := gin.Default()
	grp1 := r1.Group("/product-api")
	{
		grp1.GET("product", Controllers.GetProducts)
		grp1.POST("product", Controllers.CreateProduct)
		grp1.PATCH("product/:id", Controllers.UpdateProduct)
		grp1.GET("product/:id", Controllers.GetProductByID)
		grp1.DELETE("product/:id", Controllers.DeleteProduct)
	}

	return r1
}
func SetupRouterOrder() *gin.Engine {
	r2:=gin.Default()
	grp2:=r2.Group("/order-api")
	{
		grp2.POST("order",Controllers.CreateOrder)
		grp2.GET("order/:id",Controllers.GetOrderByID)
	}
	return r2
}
