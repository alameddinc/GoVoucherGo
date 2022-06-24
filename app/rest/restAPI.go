package main

import (
	"log"
	"net/http"
	"time"

	"github.com/alameddinc/GoVoucherGo/internal/core/handler/restApi"
	"github.com/alameddinc/GoVoucherGo/internal/core/services"
	"github.com/alameddinc/GoVoucherGo/pkg/mocks"
	"github.com/gorilla/mux"
)

func main() {
	//router, _ := restapi.NewRestMux(mux.NewRouter())
	router := mux.NewRouter()
	orderRepo := new(mocks.MockOrderRepository)
	voucherRepo := new(mocks.MockVoucherRepository)
	voucherService := services.NewVoucherService(orderRepo, voucherRepo)
	restApi.NewVoucherRestHandler(voucherService, router)
	srv := &http.Server{
		Handler:      router,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	log.Printf("Listening on port 8000")
	log.Fatal(srv.ListenAndServe())
}
