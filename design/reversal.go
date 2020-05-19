package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Note the TransactionID. This is the transaction ID of the reversal
// request itself, not the original transaction which was being reversed.
// The reversal request itself gets its own transaction ID.
var ReversalResult = Type("ReversalResult", func() {
	Description("Successful Reversal Callback")
	Attribute("Result", func() {
		Attribute("ResultType", Int, func() {
			Example(0)
		})
		Attribute("ResultCode", Int, func() {
			Example(0)
		})
		Attribute("ResultDesc", String, func() {
			Example("The service request has been accepted successfully.")
		})
		Attribute("OriginatorConversationID", String, func() {
			Example("10819-695089-1")
		})
		Attribute("ConversationID", String, func() {
			Example("AG_20170727_00004efadacd98a01d15")
		})
		Attribute("TransactionID", String, func() {
			Example("LGR019G3J2")
		})
		Attribute("ReferenceData", func() {
			Attribute("ReferenceItem", MapOf(String, String), func() {
				Key(func() {
					MinLength(1)
					Example("QueueTimeoutURL")
				})
				Elem(func() {
					Pattern("[a-zA-Z]+")
					Example("https://internalsandbox.safaricom.co.ke/mpesa/reversalresults/v1/submit")
				})
			})
		})
	})
})
