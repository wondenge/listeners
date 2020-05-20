package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// When a transaction is complete, Jenga  calls back this specified
// endpoint by making an HTTPS POST request.
var _ = Service("jenga", func() {

	HTTP(func() {

		Path("/equitybank")
	})

	// If callback URL is not reachable or returns a non-successful response,
	// Equity Bank will not retry POSTing the result of the transaction.
	Method("publish", func() {
		Description("Equity Bank Mpesa Async Callback")

		Payload(JengaCallbackPayload)

		// Endpoint returns a 200 HTTP code, and this payload.
		// Otherwise, the callback will be considered a failure.
		// No retries have been enabled at the moment.
		Result(JengaCallbackMedia)
		HTTP(func() {

			// Implementing Basic Authentication for the configured callback URL.
			// Authorization: Basic 67ew8n31me
			Header("auth:Authorization", String, "Auth token", func() {
				Pattern("^Bearer [^ ]+$")
			})

			POST("/mpesa")

			Response(StatusOK)
		})
	})

	Method("alerts", func() {
		Description("Account Alerts of monies in and out of your Equity Bank account")
		Payload(AccountAlerts)

		Result(String)
		HTTP(func() {

			// Implementing Basic Authentication for the configured callback URL.
			// Authorization: Basic 67ew8n31me
			Header("auth:Authorization", String, "Auth token", func() {
				Pattern("^Bearer [^ ]+$")
			})

			POST("/alerts")

			Response(StatusCreated)
		})
	})
})

var JengaCallbackPayload = Type("JengaCallbackPayload", func() {
	Description("Mpesa Async Callback")

	Attribute("auth", String, "Basic Authentication")

	Attribute("mobileNumber", String, func() {
		Description("recipient phone number")
		Example("254796778039")
	})
	Attribute("message", String, func() {
		Description("recipient name")
		Example("254796778039 Peter Test ")
	})
	Attribute("rrn", String, func() {
		Description("reference number")
		Example(" 190508120533 ")
	})
	Attribute("service", String, func() {
		Description("for M-Pesa the value is M-Pesa")
		Example("Mpesa")
	})
	Attribute("tranID", String, func() {
		Description("M-Pesa receipt number")
		Example(" NC84PQMWGZ")
	})
	Attribute("resultCode", String, func() {
		Description("M-Pesa return code")
		Example("0")
	})
	Attribute("resultDescription", String, func() {
		Description("M-Pesa return code description")
		Example("The service request is processed successfully.")
	})
	Attribute("additionalInfo", String, func() {
		Description("Additional information")
		Example("Additional information")
	})
})

var JengaCallbackMedia = Type("JengaCallbackMedia", func() {
	Attribute("Code", String, func() {
		Example("0")
	})
	Attribute("Message", String, func() {
		Example("OK")
	})
})

var AccountAlerts = Type("AccountAlerts", func() {

	Attribute("auth", String, "Basic Authentication")
	Extend(Customer)
	Extend(Transaction)
	Extend(Bank)
})

var Customer = Type("Customer", func() {
	Attribute("name", String, func() {
		Example("A N Other")
	})
	Attribute("mobileNumber", String, func() {
		Example("254796778039")
	})
	Attribute("reference", String)
})

var Bank = Type("Bank", func() {
	Attribute("reference", String, func() {
		Description(" S2596405")
	})
	Attribute("transactionType", String, func() {
		Description("transactionType")
		Example("C")
	})
	Attribute("account", String, func() {
		Description("account")
		Example("0111234241028")
	})
})

var Transaction = Type("Transaction", func() {

	Attribute("date", String, func() {
		Description("date")
		Format(FormatDateTime)
		Example("2018-11-27 00:00:00.0")
	})
	Attribute("reference", String, func() {
		Description("reference")
		Example(" S2596405")
	})
	Attribute("paymentMode", String, func() {
		Description("paymentMode")
		Example("TPG")
	})
	Attribute("amount", String, func() {
		Description("amount")
		Example("10")
	})
	Attribute("till", String)
	Attribute("billNumber", String, func() {
		Description("billNumber")
		Example("A N Other")
	})
	Attribute("orderAmount", String, func() {
		Description("orderAmount")
		Example("10")
	})
	Attribute("serviceCharge", String, func() {
		Description("serviceCharge")
		Example("10")
	})
	Attribute("servedBy", String, func() {
		Example("EQ")
	})
	Attribute("additionalInfo", String, func() {
		Description("additionalInfo")
		Example("MPS 254723000000 MKR35QEKV7 A N Other/537620")
	})
})
