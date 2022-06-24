package voucher

import (
	"github.com/alameddinc/GoVoucherGo/internal/core/entity"
)

type Repository interface {
	Save(entity.Voucher) error
	Update(entity.Voucher) error
	Delete(voucherCode, vendorID string) error
	Validate(voucherCode, vendorID string, cost float64) error
	Apply(voucherCode, orderId, vendorID string, cost float64) error
	Cancel(orderID string, cost float64) error
}
