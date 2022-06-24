package main

import (
	"github.com/alameddinc/GoVoucherGo/internal/core/handler/restApi"
	"github.com/alameddinc/GoVoucherGo/internal/core/services"
	"github.com/alameddinc/GoVoucherGo/pkg/mocks"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"time"
)

func main() {
	//router, _ := restapi.NewRestMux(mux.NewRouter())
	router := mux.NewRouter()
	orderRepo := new(mocks.MockOrderRepository)
	voucherRepo := new(mocks.MockVoucherRepository)
	voucherService := services.NewVoucherService(orderRepo, voucherRepo)
	restApi.NewVoucherRestHandler(voucherService, router)
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
