package handler

import "github.com/alameddinc/GoVoucherGo/internal/core/domains/voucher"

type Handler interface {
	CreateVoucher(request voucher.Request) voucher.Response
	UpdateVoucher(request voucher.Request) voucher.Response
	ValidateVoucher(request voucher.Request) voucher.Response
	ApplyOrder(request voucher.Request) voucher.Response
	CancelOrder(request voucher.Request) voucher.Response
	GetByOrder(request voucher.Request) voucher.Response
}
