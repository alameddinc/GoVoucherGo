package entity

import (
	"errors"
	"time"
)

type Order struct {
	OrderID        string
	VendorID       string
	ParentVendorID string
	TotalCost      float64
	DiscountedCost float64
	VoucherCode    string
	CreatedAt      time.Time
	CancelledAt    time.Time
}

func NewOrder(orderID, VoucherCode, vendorID, parentVendorID string, totalCost float64) (*Order, error) {
	if totalCost <= 0 {
		return nil, errors.New("order cost has to be bigger then zero")
	}
	return &Order{
		OrderID:        orderID,
		VendorID:       vendorID,
		ParentVendorID: parentVendorID,
		VoucherCode:    VoucherCode,
		TotalCost:      totalCost,
		CreatedAt:      time.Now(),
	}, nil
}

func (o *Order) ApplyDiscount(voucher Voucher) error {
	if o.VendorID != voucher.VendorID || o.ParentVendorID != voucher.VendorID {
		return errors.New("this voucher is not valid")
	}
	discount := voucher.Discount
	if voucher.Type == RateVoucher {
		discount = voucher.Discount * o.TotalCost
	}
	if discount < voucher.MinimumCost {
		return errors.New("your order's cost is lowwer then this voucher minimum amount")
	}
	o.DiscountedCost = o.TotalCost - discount
	return nil
}

func (o *Order) Cancelled() {
	o.CancelledAt = time.Now()
}
