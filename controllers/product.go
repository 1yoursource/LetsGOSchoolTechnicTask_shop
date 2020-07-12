package controllers

import (
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strconv"
	"webServer/models"

	"github.com/gin-gonic/gin"
)

var redirectCode = 301

var productModel = new(models.ProductModel)

type ProductController struct{}

//сделать 1 метод для этих функций
func (p *ProductController) GetAllProducts(c *gin.Context) {
	products := GetAllProducts(c)
	c.HTML(http.StatusOK, "products.html", products)

}
func (p *ProductController) AuthGetAllProducts(c *gin.Context) {
	products := GetAllProducts(c)
	baskets := GetUserBaskets(c)
	//c.JSON(http.StatusOK,result)
	c.HTML(http.StatusOK, "products.html", bson.M{"product": products, "basket": baskets})

}

//сделать 1 метод для этих функций
//для поиска продуктов по имени
func (p *ProductController) GetProductByName(c *gin.Context) {

	name, ok := c.GetQuery("name")

	if !ok {
		c.JSON(404, gin.H{"message": "Not received"})
	}
	if c.BindJSON(name) == nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields!"})
		c.Abort()
		return
	}

	product, err := productModel.GetProductByName(name)
	if err != nil {
		c.JSON(400, gin.H{"message": "Something problem, try again later"})
		c.Abort()
		return
	}
	//c.JSON(201, gin.H{"message": "This is your product","Product":result})
	c.HTML(200, "products.html", product)
}
func (p *ProductController) GetProductByCategory(c *gin.Context) {
	category := c.Param("category")
	result, err := productModel.GetProductByCategory(category)
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem with receiving products by category"})
		c.Abort()
		return
	}
	c.HTML(200, "products.html", result)
	//c.JSON(201, gin.H{"message": "This is your products:","Product":result})
}
func (p *ProductController) SortProductsByPrice(c *gin.Context) {
	value := c.Param("value")
	//c.Get()
	products, err := productModel.GetAllProducts()
	if err != nil {
		c.JSON(400, gin.H{"message": "Can`t get all products!"})
	}
	tempFlag, _ := strconv.ParseBool(value)
	//maybe add check parse error
	sortingProducts := SortProductsByPrice(products, tempFlag)
	c.JSON(http.StatusOK, gin.H{"message": sortingProducts})
}

func (p *ProductController) GetProductsForMainPage(c *gin.Context) {

	var productsOnMainPage = [4]string{"Pineapple Premium", "Cruasane Panavi", "Juice Sandora white grape", "Coconut Milk"}
	var products []models.Product
	for i := 0; i < len(productsOnMainPage); i++ {
		result, err := productModel.GetProductByName(productsOnMainPage[i])
		if err != nil {
			c.JSON(400, gin.H{"message": err.Error()})
			return
		}
		products = append(products, result)
	}
	//c.JSON(http.StatusOK,result)
	c.HTML(http.StatusOK, "index.html", products)
}
