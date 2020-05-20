package main

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	kithttp "github.com/go-kit/kit/transport/http"
	health "github.com/wondenge/listeners/gen/health"
	healthkitsvr "github.com/wondenge/listeners/gen/http/health/kitserver"
	healthsvr "github.com/wondenge/listeners/gen/http/health/server"
	jengakitsvr "github.com/wondenge/listeners/gen/http/jenga/kitserver"
	jengasvr "github.com/wondenge/listeners/gen/http/jenga/server"
	mpesakitsvr "github.com/wondenge/listeners/gen/http/mpesa/kitserver"
	mpesasvr "github.com/wondenge/listeners/gen/http/mpesa/server"
	swaggerkitsvr "github.com/wondenge/listeners/gen/http/swagger/kitserver"
	swaggersvr "github.com/wondenge/listeners/gen/http/swagger/server"
	jenga "github.com/wondenge/listeners/gen/jenga"
	mpesa "github.com/wondenge/listeners/gen/mpesa"
	goahttp "goa.design/goa/v3/http"
	httpmdlwr "goa.design/goa/v3/http/middleware"
	"goa.design/goa/v3/middleware"
)

// handleHTTPServer starts configures and starts a HTTP server on the given
// URL. It shuts down the server if any error is received in the error channel.
func handleHTTPServer(ctx context.Context, u *url.URL, healthEndpoints *health.Endpoints, mpesaEndpoints *mpesa.Endpoints, jengaEndpoints *jenga.Endpoints, wg *sync.WaitGroup, errc chan error, logger log.Logger, debug bool) {

	// Provide the transport specific request decoder and response encoder.
	// The goa http package has built-in support for JSON, XML and gob.
	// Other encodings can be used by providing the corresponding functions,
	// see goa.design/implement/encoding.
	var (
		dec = goahttp.RequestDecoder
		enc = goahttp.ResponseEncoder
	)

	// Build the service HTTP request multiplexer and configure it to serve
	// HTTP requests to the service endpoints.
	var mux goahttp.Muxer
	{
		mux = goahttp.NewMuxer()
	}

	// Wrap the endpoints with the transport specific layers. The generated
	// server packages contains code generated from the design which maps
	// the service input and output data structures to HTTP requests and
	// responses.
	var (
		swaggerServer                               *swaggersvr.Server
		healthShowHandler                           *kithttp.Server
		healthServer                                *healthsvr.Server
		mpesaAccountBalanceTimeoutHandler           *kithttp.Server
		mpesaAccountBalanceResultEndpointHandler    *kithttp.Server
		mpesaTransactionStatusTimeoutHandler        *kithttp.Server
		mpesaTransactionStatusResultEndpointHandler *kithttp.Server
		mpesaReversalTimeoutHandler                 *kithttp.Server
		mpesaReversalResultEndpointHandler          *kithttp.Server
		mpesaB2CTimeoutHandler                      *kithttp.Server
		mpesaB2CResultHandler                       *kithttp.Server
		mpesaC2BValidationHandler                   *kithttp.Server
		mpesaC2BConfirmationHandler                 *kithttp.Server
		mpesaServer                                 *mpesasvr.Server
		jengaPublishHandler                         *kithttp.Server
		jengaAlertsHandler                          *kithttp.Server
		jengaServer                                 *jengasvr.Server
	)
	{
		eh := errorHandler(logger)
		swaggerServer = swaggersvr.New(nil, mux, dec, enc, eh, nil)
		healthShowHandler = kithttp.NewServer(
			endpoint.Endpoint(healthEndpoints.Show),
			func(context.Context, *http.Request) (request interface{}, err error) { return nil, nil },
			healthkitsvr.EncodeShowResponse(enc),
		)
		healthServer = healthsvr.New(healthEndpoints, mux, dec, enc, eh, nil)
		mpesaAccountBalanceTimeoutHandler = kithttp.NewServer(
			endpoint.Endpoint(mpesaEndpoints.AccountBalanceTimeout),
			mpesakitsvr.DecodeAccountBalanceTimeoutRequest(mux, dec),
			mpesakitsvr.EncodeAccountBalanceTimeoutResponse(enc),
		)
		mpesaAccountBalanceResultEndpointHandler = kithttp.NewServer(
			endpoint.Endpoint(mpesaEndpoints.AccountBalanceResultEndpoint),
			mpesakitsvr.DecodeAccountBalanceResultEndpointRequest(mux, dec),
			mpesakitsvr.EncodeAccountBalanceResultEndpointResponse(enc),
		)
		mpesaTransactionStatusTimeoutHandler = kithttp.NewServer(
			endpoint.Endpoint(mpesaEndpoints.TransactionStatusTimeout),
			mpesakitsvr.DecodeTransactionStatusTimeoutRequest(mux, dec),
			mpesakitsvr.EncodeTransactionStatusTimeoutResponse(enc),
		)
		mpesaTransactionStatusResultEndpointHandler = kithttp.NewServer(
			endpoint.Endpoint(mpesaEndpoints.TransactionStatusResultEndpoint),
			mpesakitsvr.DecodeTransactionStatusResultEndpointRequest(mux, dec),
			mpesakitsvr.EncodeTransactionStatusResultEndpointResponse(enc),
		)
		mpesaReversalTimeoutHandler = kithttp.NewServer(
			endpoint.Endpoint(mpesaEndpoints.ReversalTimeout),
			mpesakitsvr.DecodeReversalTimeoutRequest(mux, dec),
			mpesakitsvr.EncodeReversalTimeoutResponse(enc),
		)
		mpesaReversalResultEndpointHandler = kithttp.NewServer(
			endpoint.Endpoint(mpesaEndpoints.ReversalResultEndpoint),
			mpesakitsvr.DecodeReversalResultEndpointRequest(mux, dec),
			mpesakitsvr.EncodeReversalResultEndpointResponse(enc),
		)
		mpesaB2CTimeoutHandler = kithttp.NewServer(
			endpoint.Endpoint(mpesaEndpoints.B2CTimeout),
			mpesakitsvr.DecodeB2CTimeoutRequest(mux, dec),
			mpesakitsvr.EncodeB2CTimeoutResponse(enc),
		)
		mpesaB2CResultHandler = kithttp.NewServer(
			endpoint.Endpoint(mpesaEndpoints.B2CResult),
			mpesakitsvr.DecodeB2CResultRequest(mux, dec),
			mpesakitsvr.EncodeB2CResultResponse(enc),
		)
		mpesaC2BValidationHandler = kithttp.NewServer(
			endpoint.Endpoint(mpesaEndpoints.C2BValidation),
			mpesakitsvr.DecodeC2BValidationRequest(mux, dec),
			mpesakitsvr.EncodeC2BValidationResponse(enc),
		)
		mpesaC2BConfirmationHandler = kithttp.NewServer(
			endpoint.Endpoint(mpesaEndpoints.C2BConfirmation),
			mpesakitsvr.DecodeC2BConfirmationRequest(mux, dec),
			mpesakitsvr.EncodeC2BConfirmationResponse(enc),
		)
		mpesaServer = mpesasvr.New(mpesaEndpoints, mux, dec, enc, eh, nil)
		jengaPublishHandler = kithttp.NewServer(
			endpoint.Endpoint(jengaEndpoints.Publish),
			jengakitsvr.DecodePublishRequest(mux, dec),
			jengakitsvr.EncodePublishResponse(enc),
		)
		jengaAlertsHandler = kithttp.NewServer(
			endpoint.Endpoint(jengaEndpoints.Alerts),
			jengakitsvr.DecodeAlertsRequest(mux, dec),
			jengakitsvr.EncodeAlertsResponse(enc),
		)
		jengaServer = jengasvr.New(jengaEndpoints, mux, dec, enc, eh, nil)
	}

	// Configure the mux.
	swaggerkitsvr.MountGenHTTPOpenapiJSON(mux)
	healthkitsvr.MountShowHandler(mux, healthShowHandler)
	mpesakitsvr.MountAccountBalanceTimeoutHandler(mux, mpesaAccountBalanceTimeoutHandler)
	mpesakitsvr.MountAccountBalanceResultEndpointHandler(mux, mpesaAccountBalanceResultEndpointHandler)
	mpesakitsvr.MountTransactionStatusTimeoutHandler(mux, mpesaTransactionStatusTimeoutHandler)
	mpesakitsvr.MountTransactionStatusResultEndpointHandler(mux, mpesaTransactionStatusResultEndpointHandler)
	mpesakitsvr.MountReversalTimeoutHandler(mux, mpesaReversalTimeoutHandler)
	mpesakitsvr.MountReversalResultEndpointHandler(mux, mpesaReversalResultEndpointHandler)
	mpesakitsvr.MountB2CTimeoutHandler(mux, mpesaB2CTimeoutHandler)
	mpesakitsvr.MountB2CResultHandler(mux, mpesaB2CResultHandler)
	mpesakitsvr.MountC2BValidationHandler(mux, mpesaC2BValidationHandler)
	mpesakitsvr.MountC2BConfirmationHandler(mux, mpesaC2BConfirmationHandler)
	jengakitsvr.MountPublishHandler(mux, jengaPublishHandler)
	jengakitsvr.MountAlertsHandler(mux, jengaAlertsHandler)

	// Wrap the multiplexer with additional middlewares. Middlewares mounted
	// here apply to all the service endpoints.
	var handler http.Handler = mux
	{
		handler = httpmdlwr.Log(logger)(handler)
		handler = httpmdlwr.RequestID()(handler)
	}

	// Start HTTP server using default configuration, change the code to
	// configure the server as required by your service.
	srv := &http.Server{Addr: u.Host, Handler: handler}
	for _, m := range swaggerServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range healthServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range mpesaServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}
	for _, m := range jengaServer.Mounts {
		logger.Log("info", fmt.Sprintf("HTTP %q mounted on %s %s", m.Method, m.Verb, m.Pattern))
	}

	(*wg).Add(1)
	go func() {
		defer (*wg).Done()

		// Start HTTP server in a separate goroutine.
		go func() {
			logger.Log("info", fmt.Sprintf("HTTP server listening on %q", u.Host))
			errc <- srv.ListenAndServe()
		}()

		<-ctx.Done()
		logger.Log("info", fmt.Sprintf("shutting down HTTP server at %q", u.Host))

		// Shutdown gracefully with a 30s timeout.
		ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
		defer cancel()

		srv.Shutdown(ctx)
	}()
}

// errorHandler returns a function that writes and logs the given error.
// The function also writes and logs the error unique ID so that it's possible
// to correlate.
func errorHandler(logger log.Logger) func(context.Context, http.ResponseWriter, error) {
	return func(ctx context.Context, w http.ResponseWriter, err error) {
		id := ctx.Value(middleware.RequestIDKey).(string)
		w.Write([]byte("[" + id + "] encoding: " + err.Error()))
		logger.Log("info", fmt.Sprintf("[%s] ERROR: %s", id, err.Error()))
	}
}
