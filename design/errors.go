package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var ErrorResponseMessage = Type("ErrorResponseMessage", func() {
	Description("Response Error Details")

	Attribute("requestId", String, func() {
		Description("This is a unique requestID for the payment request.")
		Example("16813-15-1 ")
	})

	// This are defined in the Response Error Details below.
	// The error codes maps to specific error message as illustrated
	// in the Response Error Details below.
	Attribute("errorCode", String, func() {
		Description("This is a predefined code that indicates the reason for request failure.")
		Example("404.001.04")
	})
	Attribute("errorMessage", String, func() {
		Description("This is a short descriptive message of the failure reason.")
		Example("Invalid Access Token")
	})
})

// Transaction Status API
// Response Error Details
// 404.001.04 Invalid Authentication Header
// 400.002.02 Bad Request
// 400.002.05 Invalid Request Payload
// 500.001.1001 Server Error
// 404.001.01 Resource not found
// 404.001.03 Invalid Access Token

// Transaction Reversal API
// Response Error Details
// 404.001.04 Invalid Authentication Header
// 400.002.02 Bad Request
// 400.002.05 Invalid Request Payload
// 500.001.1001 Server Error
// 404.001.01 Resource not found
// 404.001.03 Invalid Access Token

// Account Balance API
// Response Error Details
// 404.001.04 Invalid Authentication Header
// 400.002.02 Bad Request
// 400.002.05 Invalid Request Payload
// 500.001.1001 Server Error
// 404.001.01 Resource not found
// 404.001.03 Invalid Access Token

// B2C Payment API
// Response Error Details
// 400.002.01 Invalid Access Token
// 400.002.02 Bad Request
// 500.002.03 Error Occurred: Quota Violation
// 500.002.1001 Server Error
// 500.002.02 Error Occurred: Spike Arrest Violation
// 404.002.01 Resource not found
// 401.002.01 Error Occurred - Invalid Access Token

// C2B APIs
// Response Error Details
// 500.003.1001 Internal Server Error
// 400.003.01 Invalid Access Token
// 400.003.02 Bad Request
// 500.003.03 Error Occurred: Quota Violation
// 500.003.1001 Server Error
// 500.003.02 Error Occurred: Spike Arrest Violation
// 404.003.01 Resource not found
// 401.003.01 Error Occurred - Invalid Access Token

// LNMO APIs.
// 404.001.04 Invalid Authentication Header
// 400.002.02 Bad Request
// 400.002.05 Invalid Request Payload
// 500.001.1001 Server Error
// 404.001.01 Resource not found
// 404.001.03 Invalid Access Token

// Generate Token APIs
// Response Error Details
// 400.008.02 Invalid grant type passed
// 400.008.01 Invalid Authentication passed

// 		// Headers describes HTTP request/response headers.
//		Headers(func() {
//			Header("Authorization", String, func() {
//				Description("Auth2.0 verification")
//				Pattern("^Bearer [^ ]+$")
//
//				// "Bearer" keyword followed by a space
//				// and generated Access Token from OAuth API.
//				Example("Bearer xxxxxxx")
//			})
//			Header("Content-Type", String, func() {
//				Description("The requests content type.")
//				Default("application/json")
//			})
//			Required("Authorization", "Content-Type")
//		})
