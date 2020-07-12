package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"regexp"
	"sort"
	"webServer/models"
	"webServer/services"
)

func GetCurrentUser(c *gin.Context) (user models.User) {
	token, err := c.Cookie("token")
	userData := services.Tokens[token]
	user, err = UserModel.GetUserByEmail(userData.Email)
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem достать текущий юзер айди"})
		c.Abort()
		return
	}
	return user
}
func GetUserBaskets(c *gin.Context) (baskets []models.Basket) {
	user := GetCurrentUser(c)
	userWithBaskIds, err := UserModel.GetAllUserBasketsIds(user.ID)
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		return
	}

	baskets = make([]models.Basket, len(userWithBaskIds.BasketIDs), cap(userWithBaskIds.BasketIDs))
	for i := 0; i < len(userWithBaskIds.BasketIDs); i++ {
		basket, _ := BasketModel.GetBasketById(userWithBaskIds.BasketIDs[i])
		baskets[i] = basket
	}
	return baskets
}
func GetAllProducts(c *gin.Context) (products []*models.Product) {
	products, err := productModel.GetAllProducts()
	if products == nil {
		c.JSON(403, gin.H{"message": "Has no products!"})
		c.Abort()
		return
	}
	if err != nil {
		c.JSON(400, gin.H{"message": err.Error()})
		c.Abort()
		return
	}
	return products
}

/*bool=true -> min to max*/
func SortProductsByPrice(products []*models.Product, tempFlag bool) []*models.Product {
	sortingProductsPrice := Sorting(products, tempFlag)
	sortedProducts := make([]*models.Product, len(products), cap(products))
	for i := 0; i < len(sortingProductsPrice); i++ {
		for j := 0; j < len(products); j++ {
			if sortingProductsPrice[i] == products[j].Price {
				//sortedProducts = append(sortedProducts, products[i])
				sortedProducts[i] = products[j]
			}
		}
	}
	return sortedProducts
}
func Sorting(products []*models.Product, tempFlag bool) []float64 {
	sortArray := make([]float64, len(products), cap(products))
	for i := 0; i < len(products); i++ {
		//sortArray[i]+=products[i].Price
		//sortArray = append(sortArray,products[i].Price)
		sortArray[i] = products[i].Price
	}
	if tempFlag {
		sort.Float64s(sortArray)
		return sortArray
	} else {
		sort.Sort(sort.Reverse(sort.Float64Slice(sortArray)))
		return sortArray
	}
}

func ChekProductPrice(price float64) bool {
	var tempFlag bool
	if price <= 0 {
		return true
	} else {
		length := fmt.Sprintf("%f", price)
		tempPrice := make([]string, len(length))
		for i := 0; i < len(tempPrice); i++ {
			if tempPrice[i] == "e" {
				tempFlag = true
			} else {
				tempFlag = false
			}
		}
	}
	return tempFlag
}
func ChekProductCategory(category string) bool {
	return ChekHasNoDigits(category)
}
func ChekHasNoDigits(str string) bool {
	pattern := "[[:digit:]]"
	matched, _ := regexp.MatchString(pattern, str)
	return matched
}
