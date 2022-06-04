package services

import (
	"github.com/alameddinc/GoVoucherGo/internal/core/domains/order"
	"github.com/alameddinc/GoVoucherGo/internal/core/domains/voucher"
	"github.com/alameddinc/GoVoucherGo/internal/core/entity"
)

type VoucherService struct {
	orderRepo   order.Repository
	voucherRepo voucher.Repository
}

func NewVoucherService(orderRepository order.Repository, voucherRepository voucher.Repository) *VoucherService {
	return &VoucherService{
		orderRepo:   orderRepository,
		voucherRepo: voucherRepository,
	}
}

// CreateRateVoucher function
func (s *VoucherService) CreateRateVoucher(voucherType string, discountRate float64) (*entity.Voucher, error) {
	return nil, nil
}

// CreateCostVoucher function
func (s *VoucherService) CreateCostVoucher(voucherType string, discountCost float64) (*entity.Voucher, error) {
	return nil, nil
}

// UpdateVoucher function
func (s *VoucherService) UpdateVoucher(voucherCode string, voucher entity.Voucher) (*entity.Voucher, error) {
	return nil, nil
}

// ValidateVoucher function
func (s *VoucherService) ValidateVoucher(voucherCode, vendorId, parentVendorId string) error {
	return nil
}

// ApplyOrder function
func (s *VoucherService) ApplyOrder(voucherCode, vendor entity.Vendor, order entity.Order) error {
	return nil
}

// CancelOrder function
func (s *VoucherService) CancelOrder(voucherCode string, order entity.Order) error {
	return nil
}

// GetAllByVendor function
func (s *VoucherService) GetAllByVendor(vendor entity.Vendor) ([]entity.Voucher, error) {
	return nil, nil
}

// GetByOrder function
func (s *VoucherService) GetByOrder(order entity.Order) (*entity.Voucher, error) {
	return nil, nil
}

// GetByVoucherCode function
func (s *VoucherService) GetByVoucherCode(voucher entity.Voucher) (*entity.Voucher, error) {
	return nil, nil
}
