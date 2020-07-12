package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"webServer/models"
)

var orderStatus = [4]string{"New", "Accepted", "Completed", "Denied"}

var orderModel = new(models.OrderModel)

type OrderController struct{}

//add js for all
/*func (o *OrderController) CreateNewOrder(c *gin.Context) {
	var data models.Order

	data.Name = "Order #"+data.ID.String()
	data.Description = c.PostForm("description")
	data.Status = orderStatus[0]

	if c.BindJSON(&data) == nil {
		c.JSON(406, gin.H{"message": "Provide relevant fields!"})
		c.Abort()
		return
	}
	err := orderModel.CreateNewOrder(data)
	if err != nil {
		c.JSON(400, gin.H{"message": "Problem creating an order"})
		c.Abort()
		return
	}
	//c.JSON(201, gin.H{"message": "Order was create"})
	c.Redirect(redirectCode,"/createorder")
}
*/
func (o *OrderController) ChangeOrderStatus(c *gin.Context) {
	var data models.Order

	data.Description = c.PostForm("description")
	data.Status = c.PostForm("status")

	if data.Status == "" {
		c.JSON(406, gin.H{"message": "Provide relevant fields!"})
		c.Abort()
		return
	}

	currentStatus, curStatErr := orderModel.GetOrderStatus(data)
	if curStatErr != nil {
		c.JSON(400, gin.H{"message": "Problem with order status received"})
		c.Abort()
		return
	}

	if ChekOrderStatus(currentStatus.Status, data.Status) {
		c.JSON(400, gin.H{"message": "Order status can`t contains digits and new order status can`t equal to current status!"})
		c.Abort()
	}

	err := orderModel.ChangeOrderStatus(data)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem changing an order status"})
		return
	}
	//c.JSON(201, gin.H{"message": "order status was changed"})
	//c.Redirect(301,"/changeorderstatus")
}

func (o *OrderController) CancelOrder(c *gin.Context) {
	id, ok := c.GetQuery("id") //написать js

	if !ok {
		c.JSON(404, gin.H{"message": "Not canceled"})
		//c.String(404, "Not deleted!")
	}

	status, ok := c.GetQuery("status") //написать js
	if !ok {
		c.JSON(404, gin.H{"message": "Not canceled"})
		//c.String(404, "Not deleted!")
	}
	if status != orderStatus[0] {
		c.JSON(406, gin.H{"message": "Order in processing, you can`t cancel it"})
		c.Abort()
		return
	}

	err := orderModel.DeleteOrder(id)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem canceled an order"})
		c.Abort()
		return
	}
	c.JSON(201, gin.H{"message": "Order was canceled"})
	//c.Redirect(redirectCode,"/cancelorder")
}

func (o *OrderController) DeleteOrder(c *gin.Context) {
	id, ok := c.GetQuery("id") //написать js

	if !ok {
		c.JSON(404, gin.H{"message": "Not deleted"})
		//c.String(404, "Not deleted!")
	}
	err := orderModel.DeleteOrder(id)

	if err != nil {
		c.JSON(400, gin.H{"message": "Problem deleting an order"})
		c.Abort()
		return
	}
	c.JSON(201, gin.H{"message": "Order was deleted"})
}

func (o *OrderController) AdminGetAllOrders(c *gin.Context) {

	orders, err := orderModel.GetAllOrders()
	if orders == nil {
		c.JSON(406, gin.H{"message": "Has no orders!"})
		c.Abort()
		return
	}
	if err != nil {
		c.JSON(400, gin.H{"message": "Problemr"})
		c.Abort()
		return
	}
	c.HTML(http.StatusOK, "admin_orders.html", orders)
}

func ChekOrderStatus(currentStatus, newStatus string) bool {
	if ChekHasNoDigits(newStatus) {
		return true
	} else {
		var tempFlag bool
		for i := 0; i < len(orderStatus); i++ {
			if currentStatus == orderStatus[i] && (newStatus == orderStatus[i+1] || newStatus == orderStatus[len(orderStatus)-1]) {
				tempFlag = false
			} else {
				tempFlag = true
			}
		}
		return tempFlag
	}
}
