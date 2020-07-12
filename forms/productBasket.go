package forms

type ChoosingBasketCommand struct {
	Name string `json:"name" binding:"required"`
}
type AddBasketCommand struct {
	Name string `json:"name" bson:"name"`
}
