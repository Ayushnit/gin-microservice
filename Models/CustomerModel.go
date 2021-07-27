package Models

type Customer struct {
	CustomerID int16 `json:"customer_id" gorm:"primary_key"`
}
func (b *Customer) TableName() string {
	return "customer"
}
