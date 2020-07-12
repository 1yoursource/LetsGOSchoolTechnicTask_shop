package forms

type AddProductCommand struct {
	//ID			uint64 `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string  `json:"name" binding:"required"`
	Price       float64 `json:"price" binding:"required"`
	Description string  `json:"description" binding:"required"`
	Category    string  `json:"category" binding:"required"`
	//Image image.Image	`json:"image" bson:"image"`
}
