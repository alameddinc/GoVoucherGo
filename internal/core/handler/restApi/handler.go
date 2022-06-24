package restApi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/alameddinc/GoVoucherGo/internal/core/domain/voucher"
	"github.com/alameddinc/GoVoucherGo/internal/core/entity"
	"github.com/alameddinc/GoVoucherGo/pkg"
	"github.com/gorilla/mux"
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
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, pkg.ResponseSchema{Message: err.Error()})
		return
	}
	request.Validation()
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, pkg.ResponseSchema{Message: err.Error()})
		return
	}
	fmt.Fprintf(w, "%+v", pkg.ResponseSchema{Body: voucher.VoucherCore})
	return
}

// UpdateVoucher function
func (h *VoucherRestHandler) UpdateVoucher(w http.ResponseWriter, r *http.Request) {
	request := voucher.UpdateVoucherRequest{}
	err := h.voucherService.ValidateVoucher(request.Code, request.VendorID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, pkg.ResponseSchema{Message: err.Error()})
		return
	}
	voucher, err := h.voucherService.UpdateVoucher(request.Code, entity.VoucherCore(request))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, pkg.ResponseSchema{Message: err.Error()})
		return
	}
	fmt.Fprintf(w, "%+v", pkg.ResponseSchema{Body: voucher.VoucherCore})
	return
}

// ValidateVoucher function
func (h *VoucherRestHandler) ValidateVoucher(w http.ResponseWriter, r *http.Request) {
	request := voucher.ValidateVoucherRequest{}
	err := h.voucherService.ValidateVoucher(request.VoucherCode, request.VendorID)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, pkg.ResponseSchema{Message: err.Error()})
		return
	}
	fmt.Fprintf(w, "%+v", pkg.ResponseSchema{Message: "OK"})
}

// ApplyOrder function
func (h *VoucherRestHandler) ApplyOrder(w http.ResponseWriter, r *http.Request) {
	return
}

// CancelOrder function
func (h *VoucherRestHandler) CancelOrder(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("CancelOrder Running")
	return
}

// GetByOrder function
func (h *VoucherRestHandler) GetByOrder(w http.ResponseWriter, r *http.Request) {
	return
}

func (h *VoucherRestHandler) TestMethod(writer http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	fmt.Fprintf(writer, "Test Method Running with id: %s", id)
}

func createSubRouterWithMiddleware(router *mux.Router, subPath string, f func(r *http.Request, rm *mux.RouteMatch) bool) *mux.Router {
	middleware := router.MatcherFunc(f).Subrouter()
	return middleware.PathPrefix(subPath).Subrouter()
}

func (h *VoucherRestHandler) initRouter() {
	subRouter := createSubRouterWithMiddleware(h.router, "", func(r *http.Request, rm *mux.RouteMatch) bool {
		if r.Header != nil {
			return true
		}
		return false
	})
	subRouter2 := createSubRouterWithMiddleware(subRouter, "/voucher", func(r *http.Request, rm *mux.RouteMatch) bool {
		request := voucher.CreateVoucherRequest{}
		err := json.NewDecoder(r.Body).Decode(&request)
		if err != nil {
			return false
		}
		if request.VoucherType != "araba" {
			return false
		}
		return true
	})
	subRouter2.HandleFunc("/create", h.CreateVoucher).Methods("POST")
	subRouter2.HandleFunc("/update/", h.UpdateVoucher).Methods("PUT")
	subRouter2.HandleFunc("/validate/{orderID:[a-z0-9]+}", h.ValidateVoucher).Methods("GET", "POST", "PUT", "DELETE")

	subRouter.HandleFunc("/apply", h.ApplyOrder)
	subRouter.HandleFunc("/cancel", h.CancelOrder)
	subRouter.HandleFunc("/get", h.GetByOrder).Methods("GET")
	subRouter.HandleFunc("/test/{id:[0-9]+}", h.TestMethod).Methods("GET")
}
