// Code generated by goa v3.1.2, DO NOT EDIT.
//
// mpesa client HTTP transport
//
// Command:
// $ goa gen github.com/wondenge/listeners/design

package client

import (
	"context"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	goahttp "goa.design/goa/v3/http"
)

// Client lists the mpesa service endpoint HTTP clients.
type Client struct {
	// AccountBalanceTimeout Doer is the HTTP client used to make requests to the
	// AccountBalanceTimeout endpoint.
	AccountBalanceTimeoutDoer goahttp.Doer

	// AccountBalanceResultEndpoint Doer is the HTTP client used to make requests
	// to the AccountBalanceResult endpoint.
	AccountBalanceResultEndpointDoer goahttp.Doer

	// TransactionStatusTimeout Doer is the HTTP client used to make requests to
	// the TransactionStatusTimeout endpoint.
	TransactionStatusTimeoutDoer goahttp.Doer

	// TransactionStatusResultEndpoint Doer is the HTTP client used to make
	// requests to the TransactionStatusResult endpoint.
	TransactionStatusResultEndpointDoer goahttp.Doer

	// ReversalTimeout Doer is the HTTP client used to make requests to the
	// ReversalTimeout endpoint.
	ReversalTimeoutDoer goahttp.Doer

	// ReversalResultEndpoint Doer is the HTTP client used to make requests to the
	// ReversalResult endpoint.
	ReversalResultEndpointDoer goahttp.Doer

	// B2CTimeout Doer is the HTTP client used to make requests to the B2CTimeout
	// endpoint.
	B2CTimeoutDoer goahttp.Doer

	// B2CResult Doer is the HTTP client used to make requests to the B2CResult
	// endpoint.
	B2CResultDoer goahttp.Doer

	// C2BValidation Doer is the HTTP client used to make requests to the
	// C2BValidation endpoint.
	C2BValidationDoer goahttp.Doer

	// C2BConfirmation Doer is the HTTP client used to make requests to the
	// C2BConfirmation endpoint.
	C2BConfirmationDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the mpesa service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		AccountBalanceTimeoutDoer:           doer,
		AccountBalanceResultEndpointDoer:    doer,
		TransactionStatusTimeoutDoer:        doer,
		TransactionStatusResultEndpointDoer: doer,
		ReversalTimeoutDoer:                 doer,
		ReversalResultEndpointDoer:          doer,
		B2CTimeoutDoer:                      doer,
		B2CResultDoer:                       doer,
		C2BValidationDoer:                   doer,
		C2BConfirmationDoer:                 doer,
		RestoreResponseBody:                 restoreBody,
		scheme:                              scheme,
		host:                                host,
		decoder:                             dec,
		encoder:                             enc,
	}
}

// AccountBalanceTimeout returns an endpoint that makes HTTP requests to the
// mpesa service AccountBalanceTimeout server.
func (c *Client) AccountBalanceTimeout() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeAccountBalanceTimeoutRequest(c.encoder)
		decodeResponse = DecodeAccountBalanceTimeoutResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAccountBalanceTimeoutRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AccountBalanceTimeoutDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mpesa", "AccountBalanceTimeout", err)
		}
		return decodeResponse(resp)
	}
}

// AccountBalanceResultEndpoint returns an endpoint that makes HTTP requests to
// the mpesa service AccountBalanceResult server.
func (c *Client) AccountBalanceResultEndpoint() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeAccountBalanceResultEndpointRequest(c.encoder)
		decodeResponse = DecodeAccountBalanceResultEndpointResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildAccountBalanceResultEndpointRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.AccountBalanceResultEndpointDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mpesa", "AccountBalanceResult", err)
		}
		return decodeResponse(resp)
	}
}

// TransactionStatusTimeout returns an endpoint that makes HTTP requests to the
// mpesa service TransactionStatusTimeout server.
func (c *Client) TransactionStatusTimeout() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeTransactionStatusTimeoutRequest(c.encoder)
		decodeResponse = DecodeTransactionStatusTimeoutResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildTransactionStatusTimeoutRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.TransactionStatusTimeoutDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mpesa", "TransactionStatusTimeout", err)
		}
		return decodeResponse(resp)
	}
}

// TransactionStatusResultEndpoint returns an endpoint that makes HTTP requests
// to the mpesa service TransactionStatusResult server.
func (c *Client) TransactionStatusResultEndpoint() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeTransactionStatusResultEndpointRequest(c.encoder)
		decodeResponse = DecodeTransactionStatusResultEndpointResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildTransactionStatusResultEndpointRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.TransactionStatusResultEndpointDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mpesa", "TransactionStatusResult", err)
		}
		return decodeResponse(resp)
	}
}

// ReversalTimeout returns an endpoint that makes HTTP requests to the mpesa
// service ReversalTimeout server.
func (c *Client) ReversalTimeout() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeReversalTimeoutRequest(c.encoder)
		decodeResponse = DecodeReversalTimeoutResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildReversalTimeoutRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ReversalTimeoutDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mpesa", "ReversalTimeout", err)
		}
		return decodeResponse(resp)
	}
}

// ReversalResultEndpoint returns an endpoint that makes HTTP requests to the
// mpesa service ReversalResult server.
func (c *Client) ReversalResultEndpoint() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeReversalResultEndpointRequest(c.encoder)
		decodeResponse = DecodeReversalResultEndpointResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildReversalResultEndpointRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.ReversalResultEndpointDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mpesa", "ReversalResult", err)
		}
		return decodeResponse(resp)
	}
}

// B2CTimeout returns an endpoint that makes HTTP requests to the mpesa service
// B2CTimeout server.
func (c *Client) B2CTimeout() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeB2CTimeoutRequest(c.encoder)
		decodeResponse = DecodeB2CTimeoutResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildB2CTimeoutRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.B2CTimeoutDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mpesa", "B2CTimeout", err)
		}
		return decodeResponse(resp)
	}
}

// B2CResult returns an endpoint that makes HTTP requests to the mpesa service
// B2CResult server.
func (c *Client) B2CResult() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeB2CResultRequest(c.encoder)
		decodeResponse = DecodeB2CResultResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildB2CResultRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.B2CResultDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mpesa", "B2CResult", err)
		}
		return decodeResponse(resp)
	}
}

// C2BValidation returns an endpoint that makes HTTP requests to the mpesa
// service C2BValidation server.
func (c *Client) C2BValidation() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeC2BValidationRequest(c.encoder)
		decodeResponse = DecodeC2BValidationResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildC2BValidationRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.C2BValidationDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mpesa", "C2BValidation", err)
		}
		return decodeResponse(resp)
	}
}

// C2BConfirmation returns an endpoint that makes HTTP requests to the mpesa
// service C2BConfirmation server.
func (c *Client) C2BConfirmation() endpoint.Endpoint {
	var (
		encodeRequest  = EncodeC2BConfirmationRequest(c.encoder)
		decodeResponse = DecodeC2BConfirmationResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildC2BConfirmationRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.C2BConfirmationDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("mpesa", "C2BConfirmation", err)
		}
		return decodeResponse(resp)
	}
}
