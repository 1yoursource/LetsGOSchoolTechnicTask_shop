package models

import (
	ai "github.com/night-codes/mgo-ai"
	"gopkg.in/mgo.v2"
	"webServer/forms"
	"webServer/helpers"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID        uint64   `json:"_id,omitempty" bson:"_id,omitempty"`
	Name      string   `json:"name" bson:"name"`
	Email     string   `json:"email" bson:"email"`
	Password  string   `json:"password" bson:"password"`
	BasketIDs []uint64 `json:"basket_ids" bson:"basket_ids"`
	OrderIDs  []uint64 `json:"order_ids" bson:"order_ids"`
}

type UserModel struct {
	Collection *mgo.Collection
}

var UserModelTemp = UserModel{Collection: dbConnect.Use(databaseName, "users")}

func (u *UserModel) Signup(data forms.SignupUserCommand) error {
	helpers.Counter()
	err := UserModelTemp.Collection.Insert(bson.M{
		"_id":      ai.Next("users"),
		"name":     data.Name,
		"email":    data.Email,
		"password": helpers.GeneratePasswordHash([]byte(data.Password)),
	})
	return err
}
func (u *UserModel) GetAllUsers() (users []*User, err error) {
	err = UserModelTemp.Collection.Find(bson.M{}).All(&users)
	return users, err
}
func (u *UserModel) GetUserByEmail(email string) (user User, err error) {
	err = UserModelTemp.Collection.Find(bson.M{"email": email}).One(&user)
	return user, err
}
func (u *UserModel) UpdateUserBasketIds(data User) (err error) {
	err = UserModelTemp.Collection.Update(bson.M{"_id": data.ID}, data)
	return err
}
func (u *UserModel) GetAllUserBaskets(basketIDs []uint64) (baskets []*BasketToFront, err error) {
	err = UserModelTemp.Collection.Find(bson.M{"_id": bson.M{"$in": basketIDs}}).All(&baskets)
	return baskets, err
}
func (u *UserModel) GetAllUserBasketsIds(id uint64) (user User, err error) {
	err = UserModelTemp.Collection.Find(bson.M{"_id": id}).One(&user)
	return user, err
}
func (u *UserModel) GetAllUserOrders(orderIDs []uint64) (orders []*Order, err error) {
	err = UserModelTemp.Collection.Find(bson.M{"_id": bson.M{"$in": orderIDs}}).All(&orders)
	return orders, err
}
