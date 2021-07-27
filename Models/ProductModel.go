package Models

type Product struct {
	ProductID uint `json:"product_id" gorm:"primary_key"`
	Name      string `json:"name"`
	Price     int16  `json:"price"`
	Quantity  int16  `json:"qty"`
}
func (b *Product) TableName() string {
	return "product"
}
