// Code generated by goa v3.1.2, DO NOT EDIT.
//
// health go-kit HTTP server encoders and decoders
//
// Command:
// $ goa gen github.com/wondenge/listeners/design

package server

import (
	"context"
	"net/http"

	kithttp "github.com/go-kit/kit/transport/http"
	"github.com/wondenge/listeners/gen/http/health/server"
	goahttp "goa.design/goa/v3/http"
)

// EncodeShowResponse returns a go-kit EncodeResponseFunc suitable for encoding
// health show responses.
func EncodeShowResponse(encoder func(context.Context, http.ResponseWriter) goahttp.Encoder) kithttp.EncodeResponseFunc {
	return server.EncodeShowResponse(encoder)
}
