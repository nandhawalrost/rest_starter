package controllers

import (
	"starter/initializers"
	"starter/models"

	"github.com/gin-gonic/gin"
)

func ProductsIndex(c *gin.Context) {
	var products []models.Product
	initializers.DB.Find(&products)

	c.JSON(200, gin.H{
		"products": products,
	})
}

func ProductsShow(c *gin.Context) {
	id := c.Param("id")

	var products []models.Product
	initializers.DB.First(&products, id)

	c.JSON(200, gin.H{
		"product": products,
	})
}

func ProductsSearch(c *gin.Context) {
	product_name := c.Param("product_name")
	quantity := c.Param("quantity")

	var Products []models.Product
	initializers.DB.Where("product_name LIKE ? AND quantity LIKE ?", "%"+product_name+"%", "%"+quantity+"%").Find(&Products)
	c.JSON(200, gin.H{
		"products": Products,
	})
}

func ProductsCreate(c *gin.Context) {

	var body struct {
		ProductName string
		Quantity    uint
		Active      bool
	}

	c.Bind(&body)

	product := models.Product{
		ProductName: body.ProductName,
		Quantity:    body.Quantity,
		Active:      body.Active,
	}

	result := initializers.DB.Create(&product)

	if result.Error != nil {
		c.Status(400)
		return
	}
	c.JSON(200, gin.H{
		"message": "created!",
	})
}

func ProductsUpdate(c *gin.Context) {
	id := c.Param("id")

	var body struct {
		ProductName string
		Quantity    uint
		Active      bool
	}

	c.Bind(&body)

	var business models.Product
	initializers.DB.First(&business, id)

	// IN CASE YOU DON'T NEED TO UPDATE BOOL TYPE
	// result := initializers.DB.Model(&business).Updates(models.Product{
	// 	ProductName: body.ProductName,
	// 	Quantity:    body.Quantity,
	// })

	// IN CASE YOU NEED TO UPDATE BOOL TYPE
	result := initializers.DB.Model(&business).Updates(map[string]interface{}{
		"product_name": body.ProductName, "quantity": body.Quantity, "active": body.Active,
	})

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "updated!",
	})
}

func ProductsSoftDelete(c *gin.Context) {
	id := c.Param("id")

	var products []models.Product
	result := initializers.DB.Delete(&products, id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "deleted!",
	})
}

func ProductsHardDelete(c *gin.Context) {
	id := c.Param("id")

	//IMPLEMENT RAW QUERY
	result := initializers.DB.Exec("DELETE FROM products WHERE id = ?", id)

	if result.Error != nil {
		c.Status(400)
		return
	}

	c.JSON(200, gin.H{
		"message": "deleted!",
	})
}
