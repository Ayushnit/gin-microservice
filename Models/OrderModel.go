package Models

type Order struct {
	OrderID        int16 `json:"order_id" gorm:"primary_key"`
	ProductProductID int16 `json:"product_id"`
	Quantity  int16  `json:"qty"`
	CustomerCustomerID int16 `json:"customer_id"`
	Status     string `json:"status"`
	Product Product
	Customer Customer
}
func (b *Order) TableName() string {
	return "order"
}
