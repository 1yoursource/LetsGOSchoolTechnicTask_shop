package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Default(c *gin.Context) {
	c.HTML(http.StatusOK, "reg_page.html", nil)
}
func AfterAuthorization(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func Contacts(c *gin.Context) {
	c.HTML(http.StatusOK, "contacts.html", nil)
}
func MainPage(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
func FAQPage(c *gin.Context) {
	c.HTML(http.StatusOK, "faq.html", nil)
}
func AdminOrders(c *gin.Context) {
	c.HTML(http.StatusOK, "admin_orders.html", nil)
}
func RegistrationPage(c *gin.Context) {
	c.HTML(http.StatusOK, "reg_page.html", nil)
}
