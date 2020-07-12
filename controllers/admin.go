package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"webServer/forms"
	"webServer/models"
)

func (p *ProductController) AdminGetAllProducts(c *gin.Context) {
	result, err := productModel.GetAllProducts()
	if err != nil {
		fmt.Printf(err.Error())
		c.JSON(400, gin.H{"message": "Problem ap"})
		return
	}
	c.HTML(http.StatusOK, "admin_products.html", result)
}
func (p *ProductController) AddProduct(c *gin.Context) {
	var data forms.AddProductCommand

	data.Name = c.PostForm("name")
	data.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	data.Description = c.PostForm("description")
	data.Category = c.PostForm("category")

	if data.Name == "" || fmt.Sprintf("%f", data.Price) == "" || data.Description == "" || data.Category == "" {
		c.JSON(406, gin.H{"message": "Provide relevant fields!"})
		return
	}
	if ChekProductPrice(data.Price) {
		c.JSON(402, gin.H{"message": "Price can`t be negative or zero!"})
		return
	}
	if ChekProductCategory(data.Category) {
		c.JSON(403, gin.H{"message": "Category can`t contain digits!"})
		return
	}

	result, _ := productModel.GetProductByName(data.Name)

	if result.Name != "" {
		c.JSON(405, gin.H{"message": "Product with this name is already exist!"})
		return
	}

	err := productModel.AddProduct(data)

	if err != nil {
		c.JSON(407, gin.H{"message": "Problem adding a product"})
		return
	}
	p.AdminGetAllProducts(c)
}
func (p *ProductController) ChangeProduct(c *gin.Context) {
	var data models.Product

	id := c.Param("id")
	data.ID, _ = strconv.ParseUint(id, 10, 64)
	data.Name = c.PostForm("name")
	data.Price, _ = strconv.ParseFloat(c.PostForm("price"), 64)
	data.Description = c.PostForm("description")
	data.Category = c.PostForm("category")

	if data.Name == "" || fmt.Sprintf("%f", data.Price) == "" || data.Description == "" || data.Category == "" {
		c.JSON(406, gin.H{"message": "Provide relevant fields!"})
		return
	}
	if ChekProductPrice(data.Price) {
		c.JSON(400, gin.H{"message": "Price can`t be negative or zero!"})
		return
	}

	result, _ := productModel.GetProductByName(data.Name)

	if result.Name != "" && result.Name != data.Name {
		c.JSON(403, gin.H{"message": "Product with this name is already exist!"})
		return
	}

	err := productModel.ChangeProduct(data)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem changing a product"})
		return
	}
	p.AdminGetAllProducts(c)
}
func (p *ProductController) DeleteProduct(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)

	err := productModel.DeleteProduct(id)
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem with deleting", "message1": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}

//for test
func (p *ProductController) AdminGetProductByCategory(c *gin.Context) {
	category := c.Param("category")
	result, err := productModel.GetProductByCategory(category)
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem with receiving products by category"})
		return
	}
	c.HTML(http.StatusOK, "admin_products.html", result)
}
func (p *ProductController) AdminSortProductsByPrice(c *gin.Context) {
	boolValue := c.Param("boolValue")

	products, err := productModel.GetAllProducts()
	if err != nil {
		c.JSON(400, gin.H{"message": "Can`t get all products!"})
		return
	}
	tempFlag, _ := strconv.ParseBool(boolValue) //true -> min to max
	sortingProducts := SortProductsByPrice(products, tempFlag)

	c.HTML(http.StatusOK, "admin_products.html", sortingProducts)
}
