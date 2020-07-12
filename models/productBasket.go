package models

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Basket struct {
	ID         uint64   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name       string   `json:"name" bson:"name"`
	ProductIDs []uint64 `json:"product_ids" bson:"product_ids"`
}
type BasketToFront struct {
	ID       uint64    `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string    `json:"name" bson:"name"`
	Products []Product `json:"product_ids" bson:"product_ids"`
}

type BasketModel struct {
	Collection *mgo.Collection
}

var BasketModelTemp = BasketModel{Collection: dbConnect.Use(databaseName, "baskets")}

func (b *BasketModel) AddBasket(data Basket) error {
	err := BasketModelTemp.Collection.Insert(bson.M{
		"_id":  data.ID,
		"name": data.Name,
	})
	return err
}
func (b *BasketModel) AddBasketToFront(data BasketToFront) error {
	err := BasketModelTemp.Collection.Insert(bson.M{
		"_id":  data.ID,
		"name": data.Name,
	})
	return err
}
func (b *BasketModel) DeleteBasket(id uint64) error {
	err := BasketModelTemp.Collection.Remove(bson.M{"_id": id})
	return err
}
func (b *BasketModel) GetBasketById(id uint64) (basket Basket, err error) {
	err = BasketModelTemp.Collection.Find(bson.M{"_id": id}).One(&basket)
	return basket, err
}
func (b *BasketModel) GetBasketToFrontById(id uint64) (basket BasketToFront, err error) {
	err = BasketModelTemp.Collection.Find(bson.M{"_id": id}).One(&basket)
	return basket, err
}

/*func (b *BasketModel) GetAllBasketProducts(productIDs []uint64) (baskets []*BasketToFront, err error) {
	err = UserModelTemp.Collection.Find(bson.M{"_id":bson.M{"$in":productIDs}}).All(&baskets)
	return baskets, err
}*/
/*func (b *BasketModel) GetBasketToFrontByName(name string) (basket BasketToFront, err error){
	err = BasketModelTemp.Collection.Find(bson.M{"name":name}).One(&basket)
	return basket, err
}*/
func (b *BasketModel) GetBasketByName(name string) (basket Basket, err error) {
	err = BasketModelTemp.Collection.Find(bson.M{"name": name}).One(&basket)
	return basket, err
}
func (b *BasketModel) UpdateBasketProductIds(data Basket) (err error) {
	err = BasketModelTemp.Collection.Update(bson.M{"_id": data.ID}, data)
	return err
}
func (b *BasketModel) UpdateBasketToFront(data BasketToFront) (err error) {
	err = BasketModelTemp.Collection.Update(bson.M{"_id": data.ID}, data)
	return err
}

/*func (b *BasketModel) DeleteProductFromBasket(data Basket) (err error) {
	err = UserModelTemp.Collection.Remove(bson.M{"_id": bson.M{"$in":id}})
	return err
}*/

/*
func (b *BasketModel) GetBasket(data Basket) *mgo.Collection{
	return dbConnect.Use(databaseName, "basket_"+data.ID.String())
}
func (b *BasketModel) GetBasketById(id bson.ObjectId) (baskets[] *BasketToFront,err error){
	collection:=dbConnect.Use(databaseName, "basket_"+id.String())
	err = collection.Find(bson.M{}).All(&baskets)
	return baskets,err

}
func (b *BasketModel) CreateNewBasket (data Basket) error{
	err := b.GetBasket(data).Insert(data)
	return err
}
func (b *BasketModel) DeleteBasket(data interface{})error{
	collection:= dbConnect.Use(databaseName, "basket_"+fmt.Sprintf("%v",data))
	err:= collection.Remove(data)
	return err
}
func (b *BasketModel) AddProductToBasket(basket_id,product_id interface{}) error {
	collection:= dbConnect.Use(databaseName, "basket_"+fmt.Sprintf("%v",basket_id))
	err:= collection.Insert(bson.M{"_id":bson.M{"$in":product_id}})
	return err
}
func (b *BasketModel) DeleteProductFromBasket(basket_id,product_id interface{}) error {
	collection:= dbConnect.Use(databaseName, "basket_"+fmt.Sprintf("%v",basket_id))
	err:= collection.Remove(bson.M{"_id":bson.M{"$in":product_id}})
	return err
}
func (b *BasketModel) GetProductByNameInBasket(basket_id,product_id interface{}) (product Product, err error) {
	collection:= dbConnect.Use(databaseName, "basket_"+fmt.Sprintf("%v",basket_id))
	err = collection.Find(bson.M{"_id":bson.M{"$in":product_id}}).One(&product)
	return product, err
}
func (b *BasketModel) GetAllProductsFromBasket(data interface{}) (products[] Product, err error) {
	collection := dbConnect.Use(databaseName, "basket_"+fmt.Sprintf("%v",data))
	err = collection.Find(bson.M{"_id":bson.M{}}).All(&products)
	return products, err
}
func (b *BasketModel) GetAllBaskets(data Basket) (baskets []Basket,err error) {
	err = b.GetBasket(data).Find(bson.M{}).All(&baskets)
	return baskets,err
}
*/
