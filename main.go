package main

import (
	"starter/controllers"
	"starter/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	router := gin.Default()

	// same as
	// config := cors.DefaultConfig()
	// config.AllowAllOrigins = true
	// router.Use(cors.New(config))
	router.Use(cors.Default())

	router.GET("/product", controllers.ProductsIndex)
	router.GET("/product/:id", controllers.ProductsShow)
	router.GET("/product/search/:product_name/:quantity", controllers.ProductsSearch)
	router.POST("/product", controllers.ProductsCreate)
	router.PUT("/product/:id", controllers.ProductsUpdate)
	router.DELETE("/product/softdelete/:id", controllers.ProductsSoftDelete)
	router.DELETE("/product/harddelete/:id", controllers.ProductsHardDelete)
	router.Run()
}
