package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var _ = Service("mpesa", func() {

	HTTP(func() {
		Path("/callbacks/mpesa")
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
