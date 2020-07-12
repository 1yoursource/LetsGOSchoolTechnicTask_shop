package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
	"webServer/controllers"
	"webServer/handlers"
	"webServer/middleware"
)

var r = gin.Default()

func init() {
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found")
	}
}

func main() {
	r.GET("/", controllers.Default)

	// Define controllers
	user := new(controllers.UserController)
	product := new(controllers.ProductController)
	basket := new(controllers.BasketController)
	order := new(controllers.OrderController)

	auth_group := r.Group("/auth")
	auth_group.Use(middleware.TokenLogin)

	auth_group.GET("/products", product.AuthGetAllProducts)
	auth_group.POST("/add_product_to_basket/product_id/:product_id/basket_id/:basket_id", basket.AddProductToBasket)
	//auth_group.GET("/sort_product_by_price", product.SortProductsByPrice)
	//auth_group.GET("/sort_product_by_category", product.GetProductByCategory)

	auth_group.GET("/baskets", basket.GetAllUserBaskets)
	auth_group.POST("/delete_basket/:id", basket.DeleteBasket)
	auth_group.POST("/add_basket", basket.AddBasket)
	//конкретная корзина
	auth_group.GET("/basket/:id", basket.GetAllProductsFromBasket)
	//auth_group.POST("/delete_product_from_basket/:id",basket.DeleteProductFromBasket)

	//auth_group.GET("/orders",controllers.Orders)
	//auth_group.POST("/create_order",order.CreateNewOrder)
	//auth_group.POST("/cancel_order",order.CancelOrder)
	//auth_group.POST("/delete_order",order.DeleteOrder)

	//signup endpoint
	r.POST("/signup", user.Signup)
	r.GET("/signup", controllers.RegistrationPage)
	//signin
	r.POST("/signin", user.Login)
	r.GET("/signin", controllers.AfterAuthorization)
	//
	r.GET("/ping", handlers.PingGet())
	//other endpoints
	//r.GET("/main",controllers.MainPage)
	r.GET("/main", product.GetProductsForMainPage)
	r.GET("/contacts", controllers.Contacts)
	r.GET("/faq", controllers.FAQPage)

	/**/
	adm := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		os.Getenv("ADMIN_LOGIN"): os.Getenv("ADMIN_PASSWORD"),
	}))
	/**/
	//adm:=r.Group("/admin")

	// not worked
	//adm.GET("/sort_product_by_category/:category", product.AdminGetProductByCategory)
	// 404
	//adm.GET("/sort_product_by_price/:boolValue", product.AdminSortProductsByPrice)
	//all ok
	adm.GET("/" /*,controllers.AdminProducts*/, product.AdminGetAllProducts)
	adm.POST("/delete_product/:id", product.DeleteProduct)
	adm.POST("/change_product/:id", product.ChangeProduct)
	adm.POST("/add_product", product.AddProduct)
	//add js
	adm.GET("/orders", controllers.AdminOrders, order.AdminGetAllOrders)
	adm.POST("/change_order_status", order.ChangeOrderStatus)

	r.LoadHTMLGlob("templates/*.html")
	r.Static("/static", "./static")

	r.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"message": "Not found"})
	})

	r.Use(gin.Logger())

	r.Run()
}
