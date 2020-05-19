// Code generated by goa v3.1.2, DO NOT EDIT.
//
// mpesa go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/listeners/design

package server

import (
	"net/http"

	goahttp "goa.design/goa/v3/http"
)

// MountAccountBalanceTimeoutHandler configures the mux to serve the "mpesa"
// service "AccountBalanceTimeout" endpoint.
func MountAccountBalanceTimeoutHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/mpesa/accountbalance/v1/timeout", f)
}

// MountAccountBalanceResultEndpointHandler configures the mux to serve the
// "mpesa" service "AccountBalanceResult" endpoint.
func MountAccountBalanceResultEndpointHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/mpesa/accountbalance/v1/result", f)
}

// MountTransactionStatusTimeoutHandler configures the mux to serve the "mpesa"
// service "TransactionStatusTimeout" endpoint.
func MountTransactionStatusTimeoutHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/mpesa/transactionstatus/v1/timeout", f)
}

// MountTransactionStatusResultEndpointHandler configures the mux to serve the
// "mpesa" service "TransactionStatusResult" endpoint.
func MountTransactionStatusResultEndpointHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/mpesa/transactionstatus/v1/result", f)
}

// MountReversalTimeoutHandler configures the mux to serve the "mpesa" service
// "ReversalTimeout" endpoint.
func MountReversalTimeoutHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/mpesa/reversal/v1/timeout", f)
}

// MountReversalResultEndpointHandler configures the mux to serve the "mpesa"
// service "ReversalResult" endpoint.
func MountReversalResultEndpointHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/mpesa/reversal/v1/result", f)
}

// MountB2CTimeoutHandler configures the mux to serve the "mpesa" service
// "B2CTimeout" endpoint.
func MountB2CTimeoutHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/mpesa/b2c/v1/timeout", f)
}

// MountB2CResultHandler configures the mux to serve the "mpesa" service
// "B2CResult" endpoint.
func MountB2CResultHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/mpesa/b2c/v1/result", f)
}

// MountC2BValidationHandler configures the mux to serve the "mpesa" service
// "C2BValidation" endpoint.
func MountC2BValidationHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/mpesa/c2b/v1/validation", f)
}

// MountC2BConfirmationHandler configures the mux to serve the "mpesa" service
// "C2BConfirmation" endpoint.
func MountC2BConfirmationHandler(mux goahttp.Muxer, h http.Handler) {
	f, ok := h.(http.HandlerFunc)
	if !ok {
		f = func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r)
		}
	}
	mux.Handle("POST", "/mpesa/c2b/v1/confirmation", f)
}
