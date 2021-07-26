package Models

type Product struct {
	Id        string   `json:"id" gorm:"primary_key"`
	Name      string `json:"name"`
	Price     int16  `json:"price"`
	Quantity  int16  `json:"qty"`
}
func (b *Product) TableName() string {
	return "product"
}
