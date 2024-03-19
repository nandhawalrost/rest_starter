package main

import (
	"starter/controllers"
	"starter/initializers"

	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.GET("/product", controllers.ProductsIndex)
	r.GET("/product/:id", controllers.ProductsShow)
	r.GET("/product/search/:product_name/:quantity", controllers.ProductsSearch)

	r.POST("/product", controllers.ProductsCreate)
	r.PUT("/product/:id", controllers.ProductsUpdate)
	r.DELETE("/product/softdelete/:id", controllers.ProductsSoftDelete)
	r.DELETE("/product/harddelete/:id", controllers.ProductsHardDelete)
	r.Run()
}
