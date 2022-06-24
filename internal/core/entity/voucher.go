package entity

import (
	"errors"
	"time"
)

type Voucher struct {
	VoucherCore
	CountryCode string `json:"country_code"`
}

type VoucherCore struct {
	Type        string    `json:"type"`
	Discount    float64   `json:"discount"`
	Code        string    `json:"code"`
	VendorID    string    `json:"vendor_id"`
	Count       int       `json:"count"`
	CreatedAt   time.Time `json:"created_at"`
	ExpireAt    time.Time `json:"expire_at,omitempty"`
	MinimumCost float64   `json:"minimum_cost,omitempty"`
	Status      bool      `json:"status"`
}

const (
	CostVoucher = "Cost_Voucher"
	RateVoucher = "Rate_Voucher"
)

func NewCostVoucher(code, countryCode, vendorID string, discountAmount float64) (*Voucher, error) {
	if discountAmount <= 0 {
		return nil, errors.New("discount value has to be bigger then zero")
	}
	return &Voucher{
		VoucherCore: VoucherCore{
			Type:      CostVoucher,
			Code:      code,
			VendorID:  vendorID,
			Count:     1,
			Discount:  discountAmount,
			CreatedAt: time.Now(),
		},
		CountryCode: countryCode,
	}, nil
}

func NewRateVoucher(code, countryCode, vendorID string, discountRate float64) (*Voucher, error) {
	if discountRate <= 0 || discountRate > 100 {
		return nil, errors.New("discount value has be between 0 to 100")
	}
	if discountRate < 1 {
		discountRate *= 100
	}
	return &Voucher{
		VoucherCore: VoucherCore{
			Type:      RateVoucher,
			Code:      code,
			VendorID:  vendorID,
			Count:     1,
			Discount:  discountRate,
			CreatedAt: time.Now(),
		},
		CountryCode: countryCode,
	}, nil
}

func (v *Voucher) SetCount(count int) {
	v.Count = count
}
func (v *Voucher) SetExpireAt(expireAt time.Time) {
	v.ExpireAt = expireAt
}

func (v *Voucher) SetMinimumCost(minimumCost float64) {
	v.MinimumCost = minimumCost
}
