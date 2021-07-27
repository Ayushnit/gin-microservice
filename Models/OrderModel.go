package Models

type Order struct {
	OrderID        uint `json:"order_id" gorm:"primary_key"`
	ProductProductID uint `json:"product_id"`
	Quantity  int16  `json:"qty"`
	CustomerCustomerID uint `json:"customer_id"`
	Status     string `json:"status"`
	Product Product
	Customer Customer
}
func (b *Order) TableName() string {
	return "order"
}
