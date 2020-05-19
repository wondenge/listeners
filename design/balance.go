package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var AccountBalanceResult = Type("AccountBalanceResult", func() {
	Attribute("ResultType", Int, func() {
		Example(0)
	})
	Attribute("ResultCode", Int, func() {
		Example(0)
	})
	Attribute("ResultDesc", String, func() {
		Example("The service request has b een accepted successfully.")
	})
	Attribute("OriginatorConversationID", String, func() {
		Example("19464-802673-1")
	})
	Attribute("ConversationID", String, func() {
		Example("AG_20170728_0000589b6252f7f25488")
	})
	Attribute("TransactionID", String, func() {
		Example("LGS0000000")
	})
	Attribute("ResultParameters", func() {
		Attribute("ResultParameter", func() {
			Attribute("AccountBalance", MapOf(String, String), func() {
				Key(func() {
					MinLength(1)
					Example("AccountBalance")
				})
				Elem(func() {
					Pattern("[a-zA-Z]+")
					Example("Working Account|KES|46713.00|46713.00|0.00|0.00&Float Account|KES|0.00|0.00|0.00|0.00&Utility Account|KES|49217.00|49217.00|0.00|0.00&Charges Paid Account|KES|-220.00|-220.00|0.00|0.00&Organization Settlement Account|KES|0.00|0.00|0.00|0.00")
				})

			})
			Attribute("BOCompletedTime", MapOf(String, Int), func() {
				Key(func() {
					MinLength(1)
					Example("BOCompletedTime")
				})
				Elem(func() {
					Format(FormatDateTime)
					Example(20170728095642)
				})
			})
		})
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
