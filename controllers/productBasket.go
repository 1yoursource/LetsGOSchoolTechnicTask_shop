package controllers

import (
	"github.com/gin-gonic/gin"
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"webServer/helpers"
	"webServer/models"
)

var BasketModel = new(models.BasketModel)

type BasketController struct{}

func (b *BasketController) AddBasket(c *gin.Context) {
	helpers.Counter()
	var data models.Basket
	data.Name = c.PostForm("name")
	data.ID = ai.Next("baskets")
	if data.Name == "" {
		c.JSON(400, gin.H{"message": "Provide relevants fields"})
		return
	}

	err := BasketModel.AddBasket(data)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	user := GetCurrentUser(c)

	user.BasketIDs = append(user.BasketIDs, data.ID)

	err = UserModel.UpdateUserBasketIds(user)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	b.GetAllUserBaskets(c)
}
func (b *BasketController) GetAllUserBaskets(c *gin.Context) {
	baskets := GetUserBaskets(c)
	c.HTML(200, "baskets.html", baskets)
}
func (b *BasketController) DeleteBasket(c *gin.Context) {
	id, _ := strconv.ParseUint(c.Param("id"), 10, 64)
	err := BasketModel.DeleteBasket(id)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}
	user := GetCurrentUser(c)
	for i := 0; i < len(user.BasketIDs); i++ {
		if user.BasketIDs[i] == id {
			copy(user.BasketIDs[i:], user.BasketIDs[i+1:])
			user.BasketIDs[len(user.BasketIDs)-1] = 0
			user.BasketIDs = user.BasketIDs[:len(user.BasketIDs)-1]
		}
	}
	err = UserModel.UpdateUserBasketIds(user)
	b.GetAllUserBaskets(c)
}
func (b *BasketController) AddProductToBasket(c *gin.Context) {
	product_id := c.Param("product_id") //product id
	basket_id := c.Param("basket_id")   //check its id or name

	if basket_id == "" {
		c.JSON(401, gin.H{"message": "Please, choose the basket"})
		return
	}
	basketID, _ := strconv.ParseUint(basket_id, 10, 64)
	basket, err := BasketModel.GetBasketById(basketID)
	if err != nil {
		c.JSON(402, gin.H{"message": err.Error()})
		return
	}
	productId, _ := strconv.ParseUint(product_id, 10, 64)

	basket.ProductIDs = append(basket.ProductIDs, productId)

	err = BasketModel.UpdateBasketProductIds(basket)
	if err != nil {
		c.JSON(402, gin.H{"message": err.Error()})
		return
	}
	c.Status(200)
}
func (b *BasketController) GetAllProductsFromBasket(c *gin.Context) {
	id := c.Param("id")
	basketID, _ := strconv.ParseUint(id, 10, 64)
	basketToFront, err := BasketModel.GetBasketToFrontById(basketID)
	if err != nil {
		c.JSON(401, bson.M{"message": err.Error()})
		return
	}

	c.HTML(200, "productBasket.html", basketToFront)
}
