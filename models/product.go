package models

import (
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"webServer/forms"
	"webServer/helpers"
)

type Product struct {
	ID          uint64  `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string  `json:"name" bson:"name"`
	Price       float64 `json:"price" bson:"price"`
	Description string  `json:"description" bson:"description"`
	Category    string  `json:"category" bson:"category"`
	//Image		image.Image	`json:"image" bson:"image"`
}

type ProductModel struct {
	Collection *mgo.Collection
}

var ProductModelTemp = ProductModel{Collection: dbConnect.Use(databaseName, "products")}

func (p *ProductModel) AddProduct(data forms.AddProductCommand) error {
	helpers.Counter()
	err := ProductModelTemp.Collection.Insert(bson.M{
		"_id":         ai.Next("products"),
		"name":        data.Name,
		"price":       data.Price,
		"description": data.Description,
		"category":    data.Category,
	})
	return err
}
func (p *ProductModel) ChangeProduct(data Product) error {
	err := ProductModelTemp.Collection.Update(bson.M{"_id": data.ID}, data)
	return err
}
func (p *ProductModel) DeleteProduct(id uint64) error {
	err := ProductModelTemp.Collection.Remove(bson.M{"_id": id})
	return err
}
func (p *ProductModel) GetProductByName(name string) (product Product, err error) {
	err = ProductModelTemp.Collection.Find(bson.M{"name": name}).One(&product)
	return product, err
}
func (p *ProductModel) GetProductByCategory(category string) (products []*Product, err error) {
	err = ProductModelTemp.Collection.Find(bson.M{"category": category}).All(&products)
	return products, err
}
func (p *ProductModel) GetAllProducts() (products []*Product, err error) {
	err = ProductModelTemp.Collection.Find(bson.M{}).All(&products)
	return products, err
}
func (p *ProductModel) GetProductByID(ID uint64) (product Product, err error) {
	err = ProductModelTemp.Collection.Find(bson.M{"_id": ID}).One(&product)
	return product, err
}
