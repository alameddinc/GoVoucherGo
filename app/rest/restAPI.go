package main

import (
	"github.com/alameddinc/GoVoucherGo/internal/core/domains/voucher"
	"github.com/alameddinc/GoVoucherGo/internal/core/handler/restApi"
	"github.com/alameddinc/GoVoucherGo/internal/core/services"
	"github.com/alameddinc/GoVoucherGo/pkg/mocks"
)

func main() {
	orderRepo := new(mocks.MockOrderRepository)
	voucherRepo := new(mocks.MockVoucherRepository)
	voucherService := services.NewVoucherService(orderRepo, voucherRepo)
	handler := restApi.NewVoucherRestHandler(voucherService)
	handler.ApplyOrder(voucher.Request{})
}
