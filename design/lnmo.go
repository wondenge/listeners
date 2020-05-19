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
			Attribute("MerchantRequestID")
			Attribute("CheckoutRequestID")
			Attribute("ResultCode")
			Attribute("ResultDesc")

			// This is the JSON object that holds more details
			// for the transaction. It is only returned for
			// Successful transaction.
			Attribute("CallbackMetadata", ArrayOf(Item), func() {
				Attribute("Amount", Int64, func() {
					Description("This is the Amount that was transacted")
					Example(10500.5)
				})

				// Same value is sent to customer over
				// SMS upon successful processing.
				Attribute("MpesaReceiptNumber", String, func() {
					Description("Unique M-PESA transaction ID for the payment request")
					Example("LHG31AA5TX")
				})
				Attribute("Balance", Int64, func() {
					Description("Balance of the account for the shortcode used as partyB")
					Example(32009.9)
				})
				Attribute("TransactionDate", String, func() {
					Description("Timestamp representing transaction date time completion")
					Format(FormatDateTime)
					Example("20170827163400")
				})
				Attribute("PhoneNumber", Int, func() {
					Description("Number of the customer who made the payment.")
					Example(0722000000)
				})
			})
		})
	})
})

// This is a JSON Array, within the CallbackMetadata, that holds additional
// transaction details in JSON objects. Since this array is returned under
// theCallbackMetadata, it is only returned for Successful transaction.
var Item = Type("Item", func() {
	Attribute("Name", String)
	Attribute("Value", Int)
})
