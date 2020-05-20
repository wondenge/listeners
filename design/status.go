package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var TransactionStatusResult = Type("TransactionStatusResult", func() {
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
			Attribute("ResultParameter", ArrayOf(TransactionStatusResultParameter))
		})
		Attribute("ReferenceData", func() {
			Attribute("ReferenceItem", MapOf(String, String), func() {
				Key(func() {
					MinLength(1)
					Example("Occasion")
				})
				Elem(func() {
					Pattern("[a-zA-Z]+")
					Example("Occasion")
				})
			})
		})
	})
})

var TransactionStatusResultParameter = Type("TransactionStatusResultParameter", func() {
	Attribute("ReceiptNo", String, func() {
		Example("LGR919G2AV")
	})
	Attribute("ConversationID", String, func() {
		Example("AG_20170727_00004492b1b6d0078fbe")
	})
	Attribute("FinalisedTime", String, func() {
		Format(FormatDateTime)
		Example("20170727101415")
	})
	Attribute("Amount", Int, func() {
		Example(10)
	})
	Attribute("TransactionStatus", String, func() {
		Example("Completed")
	})
	Attribute("ReasonType", String, func() {
		Example("Salary Payment via API")
	})
	Attribute("TransactionReason", String, func() {
		Example("Transaction Reason")
	})
	Attribute("DebitPartyCharges", String, func() {
		Example("Fee For B2C Payment|KES|33.00")
	})
	Attribute("DebitAccountType", String, func() {
		Example("Utility Account")
	})
	Attribute("InitiatedTime", String, func() {
		Format(FormatDateTime)
		Example("20170727101415")
	})
	Attribute("Originator Conversation ID", String, func() {
		Example("19455-773836-1")
	})
	Attribute("CreditPartyName", String, func() {
		Example("254708374149 - John Doe")
	})
	Attribute("DebitPartyName", String, func() {
		Example("600134 - Safaricom157")
	})
})
