package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var LNMOResult = ResultType("LNMOResult", func() {
	Attribute("Body", func() {
		Attribute("StkCallback", func() {
			Attribute("MerchantRequestID", String, func() {
				Example("8555-67195-1")
			})
			Attribute("CheckoutRequestID", String, func() {
				Example("ws_CO_27072017151044001")
			})
			Attribute("ResultCode", Int, func() {
				Example(1032)
			})
			Attribute("ResultDesc", String, func() {
				Example("[STK_CB - ]Request cancelled by user")
				Example("The service request is processed successfully.")
			})

			// This is the JSON object that holds more details
			// for the transaction. It is only returned for
			// Successful transaction.
			Attribute("CallbackMetadata", ArrayOf(Item))
		})
	})
})

var Item = Type("Item", func() {
	Attribute("Amount", MapOf(String, Int), func() {
		Description("This is the Amount that was transacted")
		Key(func() {
			Example("Amount")
		})
		Elem(func() {
			Example(10500)
		})
	})

	// Same value is sent to customer over
	// SMS upon successful processing.
	Attribute("MpesaReceiptNumber", MapOf(String, String), func() {
		Description("Unique M-PESA transaction ID for the payment request")
		Key(func() {
			Example("MpesaReceiptNumber")
		})
		Elem(func() {
			Example("LHG31AA5TX")
		})
	})
	Attribute("Balance", MapOf(String, Int), func() {
		Description("Balance of the account for the shortcode used as partyB")
		Key(func() {
			Example("Balance")
		})
		Elem(func() {
			Example(32009)
		})
	})
	Attribute("TransactionDate", MapOf(String, String), func() {
		Description("Timestamp representing transaction date time completion")
		Key(func() {
			Example("TransactionDate")
		})
		Elem(func() {
			Format(FormatDateTime)
			Example("20170827163400")
		})
	})
	Attribute("PhoneNumber", MapOf(String, Int), func() {
		Description("Number of the customer who made the payment.")
		Key(func() {
			Example("PhoneNumber")
		})
		Elem(func() {
			Example(0722000000)
		})
	})
})
