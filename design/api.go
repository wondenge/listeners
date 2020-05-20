package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("listener", func() {
	Title("Listener")
	Description("Golang HTTP service listening for Callback and Timeout URLs")
	Version("1.0")
	TermsOfService("https://github.com/wondenge/mpesa-listener/blob/master/LICENSE")
	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
		URL("https://www.ondenge.me")
	})
	License(func() {
		Name("Apache License")
		URL("https://github.com/wondenge/mpesa-listener/blob/master/LICENSE")
	})
	Docs(func() {
		Description("Library Usage")
		URL("https://github.com/wondenge/mpesa-listener/blob/master/README.md")
	})
	Server("listener", func() {
		Description("listener hosts all Listener Services.")
		Services("swagger", "health", "mpesa")
		Host("local", func() {
			Description("Localhost")
			URI("http://localhost:3000")
		})
	})
})

var _ = Service("swagger", func() {
	Description("The swagger service serves the API swagger definition.")
	HTTP(func() {
		Path("/swagger")
	})

	// Defines an endpoint that serves static assets via HTTP.
	// Serve the file with relative path ../../gen/http/openapi.json
	// for requests sent to /swagger.json.
	Files("/swagger.json", "../../gen/http/openapi.json", func() {
		Description("JSON document containing the API swagger definition")
	})
})

var _ = Service("health", func() {

	// HTTP defines the HTTP transport specific
	// properties of an API, a service or a single method.
	HTTP(func() {
		Path("/health")
	})

	// Defines a single service method
	Method("show", func() {
		Description("Health check endpoint.")
		Result(String)
		HTTP(func() {
			GET("/")
			Response(func() {
				Code(StatusOK)
				ContentType("text/plain")
			})
		})
	})
})

var _ = Service("mpesa", func() {

	HTTP(func() {
		Path("/mpesa")
	})

	// Account Balance Queue TimeOut URL
	// Path that stores information of time out transaction.
	Method("AccountBalanceTimeout", func() {
		Description("Account Balance Queue TimeOut URL")
		Payload(AccountBalanceResult)
		Result(String)
		HTTP(func() {
			POST("/accountbalance/v1/timeout")
			Response(StatusCreated)
		})
	})

	// Account Balance Result URL
	// Path that stores information of transaction
	Method("AccountBalanceResult", func() {
		Description("Account Balance Result URL")
		Payload(AccountBalanceResult)
		Result(String)
		HTTP(func() {
			POST("/accountbalance/v1/result")
			Response(StatusCreated)
		})
	})

	// Transaction Status Queue TimeOut URL
	// Path that stores information of time out transaction.
	Method("TransactionStatusTimeout", func() {
		Description("Transaction Status Queue TimeOut URL")
		Payload(TransactionStatusResult)
		Result(String)
		HTTP(func() {
			POST("/transactionstatus/v1/timeout")
			Response(StatusCreated)
		})
	})

	// Transaction Status Result URL
	// Path that stores information of transaction
	Method("TransactionStatusResult", func() {
		Description("Transaction Status Result URL")
		Payload(TransactionStatusResult)
		Result(String)
		HTTP(func() {
			POST("/transactionstatus/v1/result")
			Response(StatusCreated)
		})
	})

	// Reversal Queue TimeOut URL
	// Path that stores information of time out transaction.
	Method("ReversalTimeout", func() {
		Description("Reversal Queue TimeOut URL")
		Payload(ReversalResult)
		Result(String)
		HTTP(func() {
			POST("/reversal/v1/timeout")
			Response(StatusCreated)
		})
	})

	// Reversal Result URL
	// Path that stores information of transaction
	Method("ReversalResult", func() {
		Description("Reversal Result URL")
		Payload(ReversalResult)
		Result(String)
		HTTP(func() {
			POST("/reversal/v1/result")
			Response(StatusCreated)
		})
	})

	// B2C Queue TimeOut URL
	// Path that stores information of time out transaction.
	Method("B2CTimeout", func() {
		Description("B2C Queue TimeOut URL")
		Payload(B2CPaymentResult)
		Result(String)
		HTTP(func() {
			POST("/b2c/v1/timeout")
			Response(StatusCreated)
		})
	})

	// B2C Result URL
	// Path that stores information of transaction
	Method("B2CResult", func() {
		Description("B2C Result URL")
		Payload(B2CPaymentResult)
		Result(String)
		HTTP(func() {
			POST("/b2c/v1/result")
			Response(StatusCreated)
		})
	})

	// C2B Validation URL
	Method("C2BValidation", func() {
		Description("C2B Validation URL")
		Payload(ValidationResult)
		Result(String)
		HTTP(func() {
			POST("/c2b/v1/validation")
			Response(StatusCreated)
		})
	})

	// C2B Confirmation URL
	Method("C2BConfirmation", func() {
		Description("C2B Confirmation URL")
		Payload(ConfirmationResult)
		Result(String)
		HTTP(func() {
			POST("/c2b/v1/confirmation")
			Response(StatusCreated)
		})
	})
})
