package order

import "github.com/alameddinc/GoVoucherGo/internal/core/entity"

type Repository interface {
	GetDetails(orderID string) *entity.Order
	Apply(orderCost float64, orderID, VoucherCode string)
	Cancel(orderID, VoucherCode string)
}
