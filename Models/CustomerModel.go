package Models

type Customer struct {
	CustomerID uint `json:"customer_id" gorm:"primary_key"`
	Active bool      `json:"active"`
}
func (b *Customer) TableName() string {
	return "customer"
}
