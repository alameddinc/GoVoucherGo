package voucher

import "github.com/alameddinc/GoVoucherGo/internal/core/entity"

type Service interface {
	CreateRateVoucher(vendorID string, discountRate float64) (*entity.Voucher, error)
	CreateCostVoucher(vendorID string, discountCost float64) (*entity.Voucher, error)
	UpdateVoucher(voucherCode string, voucher entity.VoucherCore) (*entity.Voucher, error)

	ValidateVoucher(voucherCode, vendorId string) error
	ApplyOrder(voucherCode, vendor entity.Vendor, order entity.Order) error
	CancelOrder(voucherCode string, order entity.Order) error

	GetAllByVendor(vendor entity.Vendor) ([]entity.Voucher, error)
	GetByOrder(order entity.Order) (*entity.Voucher, error)
	GetByVoucherCode(voucher string) (*entity.Voucher, error)
}
