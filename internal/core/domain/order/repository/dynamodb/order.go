package dynamodb

import "github.com/alameddinc/GoVoucherGo/internal/core/entity"

type OrderDynamoDB struct {
}

// GetDetails function
func (d *OrderDynamoDB) GetDetails(orderID string) *entity.Order {
	return nil
} // Apply function
func (d *OrderDynamoDB) Apply(orderCost float64, orderID, VoucherCode string) {
	return
} // Cancel function
func (d *OrderDynamoDB) Cancel(orderID, VoucherCode string) {
	return
}
