package voucher

import "github.com/alameddinc/GoVoucherGo/internal/core/entity"

type CreateVoucherRequest struct {
	VoucherType string
	Discount    float64
	VendorID    string
}

type ValidateVoucherRequest struct {
	VoucherCode string
	VendorID    string
}

type UpdateVoucherRequest entity.VoucherCore
type CreateVoucherResponse entity.VoucherCore
