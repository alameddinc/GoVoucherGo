package restapi

import (
	"errors"
	"github.com/gorilla/mux"
)

type RestMux struct {
	router *mux.Router
}

// NewRestMux is a Mux Router Generator
// Example Prefixes: "abc", "bcd", "cde" => localhost/abc/bcd/cde
func NewRestMux(router *mux.Router, prefixes ...string) (*RestMux, error) {
	if router == nil {
		return nil, errors.New("router can not be empty")
	}
	if len(prefixes) != 0 {
		for _, prefix := range prefixes {
			router = router.PathPrefix("/" + prefix).Subrouter()
		}
	}
	return &RestMux{router: router}, nil
}
