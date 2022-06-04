package restApi

import "github.com/alameddinc/GoVoucherGo/internal/core/domains/voucher"

type VoucherRestHandler struct {
	voucherService voucher.Service
}

func NewVoucherRestHandler(voucherService voucher.Service) *VoucherRestHandler {
	return &VoucherRestHandler{voucherService: voucherService}
}

// CreateVoucher function
func (h *VoucherRestHandler) CreateVoucher(request voucher.Request) voucher.Response {
	return voucher.Response{}
}

// UpdateVoucher function
func (h *VoucherRestHandler) UpdateVoucher(request voucher.Request) voucher.Response {
	return voucher.Response{}
}

// ValidateVoucher function
func (h *VoucherRestHandler) ValidateVoucher(request voucher.Request) voucher.Response {
	return voucher.Response{}
}

// ApplyOrder function
func (h *VoucherRestHandler) ApplyOrder(request voucher.Request) voucher.Response {
	return voucher.Response{}
}

// CancelOrder function
func (h *VoucherRestHandler) CancelOrder(request voucher.Request) voucher.Response {
	return voucher.Response{}
}

// GetByOrder function
func (h *VoucherRestHandler) GetByOrder(request voucher.Request) voucher.Response {
	return voucher.Response{}
}
