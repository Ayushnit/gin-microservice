package Models

type Order struct {
	Id        string `json:"id" gorm:"primary_key"`
	ProductID string `json:"product_id"`
	Quantity  int16  `json:"qty"`
}
func (b *Order) TableName() string {
	return "order"
}
