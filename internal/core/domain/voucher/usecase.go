package voucher

import "github.com/alameddinc/GoVoucherGo/internal/core/entity"

type CreateVoucherRequest struct {
	VoucherType string `validation:"null"`
	Discount    float64
	VendorID    string
}

func (r CreateVoucherRequest) Validation() {

}

type ValidateVoucherRequest struct {
	VoucherCode string
	VendorID    string
}

type UpdateVoucherRequest entity.VoucherCore
type CreateVoucherResponse entity.VoucherCore
type Request entity.VoucherCore
type Response entity.VoucherCore
