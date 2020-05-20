package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var AccountBalanceResult = Type("AccountBalanceResult", func() {
	Attribute("Result", func() {
		Attribute("ResultType", Int, func() {
			Description("Status code indicating whether transaction was already sent to your listener")
			Default(0)
			Example(0)
		})

		// 0 means success.
		// Any other code means an error occurred
		// or the transaction failed.
		Attribute("ResultCode", Int, func() {
			Description("Numeric status code indicating the status of the transaction processing")
			Example(0)
		})

		// Usually maps to a specific result code value
		Attribute("ResultDesc", String, func() {
			Description("Message from the API that gives the status of the request")
			Example("Service request is has bee accepted successfully")
			Example("Initiator information is invalid")
		})

		// returned by API proxy upon successful request submission.
		Attribute("OriginatorConversationId", String, func() {
			Description("Unique identifier for the transaction request.")
			Example("AG_2376487236_126732989KJHJKH")
		})

		// returned by the M-Pesa upon successful request submission.
		Attribute("ConversationId", String, func() {
			Description("Unique identifier for the transaction request.")
			Example("236543-276372-2")
		})
		// Same value is sent to customer over SMS upon successful processing.
		Attribute("TransactionID", String, func() {
			Description("Unique M-PESA transaction ID for every payment request.")
			Example("LHG31AA5TX")
		})
		Attribute("ResultParameters", func() {
			Attribute("ResultParameter", ArrayOf(AccountBalanceParameters))
		})
		Attribute("ReferenceData", func() {
			Attribute("ReferenceItem", MapOf(String, String), func() {
				Key(func() {
					MinLength(1)
					Example("QueueTimeoutURL")
				})
				Elem(func() {
					Pattern("[a-zA-Z]+")
					Example("https://internalsandbox.safaricom.co.ke/mpesa/abresults/v1/submit")
				})
			})
		})
	})
})

var AccountBalanceParameters = Type("AccountBalanceParameters", func() {
	Attribute("AccountBalance", MapOf(String, String), func() {
		Pattern("[a-zA-Z]+")
		Example("Working Account|KES|46713.00|46713.00|0.00|0.00&Float Account|KES|0.00|0.00|0.00|0.00&Utility Account|KES|49217.00|49217.00|0.00|0.00&Charges Paid Account|KES|-220.00|-220.00|0.00|0.00&Organization Settlement Account|KES|0.00|0.00|0.00|0.00")
	})
	Attribute("BOCompletedTime", String, func() {
		Format(FormatDateTime)
		Example("20170728095642")
	})
})
