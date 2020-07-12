package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"webServer/helpers"
)

type Order struct {
	ID          uint64 `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string `json:"name" bson:"name"`
	Description string `json:"description" bson:"description"`
	Status      string `json:"status" bson:"status"`
}

type OrderModel struct {
	Collection *mgo.Collection
}

var OrderModelTemp = OrderModel{Collection: dbConnect.Use(databaseName, "orders")}

func (o *OrderModel) CreateNewOrder(data Order) error {
	helpers.Counter()
	err := OrderModelTemp.Collection.Insert(data)
	return err
}
func (o *OrderModel) ChangeOrderStatus(data Order) error {
	err := OrderModelTemp.Collection.Update(bson.M{"_id": data.ID}, data)
	return err
}
func (o *OrderModel) DeleteOrder(data interface{}) error {
	err := OrderModelTemp.Collection.Remove(data)
	return err
}
func (o *OrderModel) GetAllOrders() (order []Order, err error) {
	err = OrderModelTemp.Collection.Find(bson.M{}).All(&order)
	return order, err
}
func (o *OrderModel) GetOrderStatus(data Order) (order Order, err error) {
	err = OrderModelTemp.Collection.Find(bson.M{"_id": data.ID}).One(&order)
	return order, err
}
