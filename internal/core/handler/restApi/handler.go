package restApi

import (
	"fmt"
	"github.com/alameddinc/GoVoucherGo/internal/core/domains/voucher"
	"github.com/alameddinc/GoVoucherGo/internal/core/entity"
	"github.com/gorilla/mux"
	"net/http"
)

type VoucherRestHandler struct {
	voucherService voucher.Service
	router         *mux.Router
}

func NewVoucherRestHandler(voucherService voucher.Service, router *mux.Router) *VoucherRestHandler {
	voucherHandler := &VoucherRestHandler{
		voucherService: voucherService,
		router:         router,
	}
	voucherHandler.initRouter()
	return voucherHandler
}

// CreateVoucher function
func (h *VoucherRestHandler) CreateVoucher(w http.ResponseWriter, r *http.Request) {
	request := voucher.CreateVoucherRequest{}
	createFunc := h.voucherService.CreateRateVoucher
	if request.VoucherType == entity.CostVoucher {
		createFunc = h.voucherService.CreateCostVoucher
	}
	voucher, err := createFunc(request.VendorID, request.Discount)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, err)
		return
	}
	fmt.Fprintf(w, "%+v", ResponseSchema{Body: voucher.VoucherCore})
	return
}

// UpdateVoucher function
func (h *VoucherRestHandler) UpdateVoucher(w http.ResponseWriter, r *http.Request) {
	request := voucher.UpdateVoucherRequest{}
	err := h.voucherService.ValidateVoucher(request.Code, request.VendorID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ResponseSchema{Message: err.Error()})
		return
	}
	voucher, err := h.voucherService.UpdateVoucher(request.Code, entity.VoucherCore(request))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ResponseSchema{Message: err.Error()})
		return
	}
	fmt.Fprintf(w, "%+v", ResponseSchema{Body: voucher.VoucherCore})
	return
}

type ResponseSchema struct {
	Message string `json:"message"`
	Body    any    `json:"body"`
}

// ValidateVoucher function
func (h *VoucherRestHandler) ValidateVoucher(w http.ResponseWriter, r *http.Request) {
	request := voucher.ValidateVoucherRequest{}
	err := h.voucherService.ValidateVoucher(request.VoucherCode, request.VendorID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, ResponseSchema{Message: err.Error()})
		return
	}
	fmt.Fprintf(w, "%+v", ResponseSchema{Message: "OK"})
}

// ApplyOrder function
func (h *VoucherRestHandler) ApplyOrder(w http.ResponseWriter, r *http.Request) {
	return
}

// CancelOrder function
func (h *VoucherRestHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	return
}

// GetByOrder function
func (h *VoucherRestHandler) GetByOrder(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *VoucherRestHandler) initRouter() {
	subRouter := h.router.PathPrefix("/voucher").Subrouter()
	subRouter.HandleFunc("/create", h.CreateVoucher)
	subRouter.HandleFunc("/update", h.UpdateVoucher)
	subRouter.HandleFunc("/validate", h.ValidateVoucher)
	subRouter.HandleFunc("/apply", h.ApplyOrder)
	subRouter.HandleFunc("/cancel", h.CancelOrder)
	subRouter.HandleFunc("/get", h.GetByOrder)
}
