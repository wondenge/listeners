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
			Example(0)
		})
		Attribute("ResultCode", Int, func() {
			Example(0)
		})
		Attribute("ResultDesc", String, func() {
			Example("The service request has been accepted successfully")
		})
		Attribute("OriginatorConversationID", String, func() {
			Example("10816-694520-2")
		})
		Attribute("ConversationID", String, func() {
			Example("AG_20170727_000059c52529a8e080bd")
		})
		Attribute("TransactionID", String, func() {
			Example("LGR0000000")
		})
		Attribute("ResultParameters", func() {
			Attribute("ResultParameter", func() {
				Attribute("ReceiptNo", MapOf(String, String), func() {
					Key(func() {
						Example("ReceiptNo")
					})
					Elem(func() {
						Example("LGR919G2AV")
					})
				})
				Attribute("Conversation ID", MapOf(String, String), func() {
					Key(func() {
						Example("Conversation ID")
					})
					Elem(func() {
						Example("AG_20170727_00004492b1b6d0078fbe")
					})
				})
				Attribute("FinalisedTime", MapOf(String, Int), func() {
					Key(func() {
						Example("FinalisedTime")
					})
					Elem(func() {
						Example(20170727101415)
					})
				})
				Attribute("Amount", MapOf(String, Int), func() {
					Key(func() {
						Example("Amount")
					})
					Elem(func() {
						Example(10)
					})
				})
				Attribute("TransactionStatus", MapOf(String, String), func() {
					Key(func() {
						Example("TransactionStatus")
					})
					Elem(func() {
						Example("Completed")
					})
				})
				Attribute("ReasonType", MapOf(String, String), func() {
					Key(func() {
						Example("ReasonType")
					})
					Elem(func() {
						Example("Salary Payment via API")
					})
				})
				Attribute("TransactionReason", MapOf(String, String), func() {
					Key(func() {
						Example("TransactionReason")
					})
					Elem(func() {
						Example("Transaction Reason")
					})
				})
				Attribute("DebitPartyCharges", MapOf(String, String), func() {
					Key(func() {
						Example("DebitPartyCharges")
					})
					Elem(func() {
						Example("Fee For B2C Payment|KES|33.00")
					})
				})
				Attribute("DebitAccountType", MapOf(String, String), func() {
					Key(func() {
						Example("DebitAccountType")
					})
					Elem(func() {
						Example("Utility Account")
					})
				})
				Attribute("InitiatedTime", MapOf(String, Int), func() {
					Key(func() {
						Example("InitiatedTime")
					})
					Elem(func() {
						Example(20170727101415)
					})
				})
				Attribute("Originator Conversation ID", MapOf(String, String), func() {
					Key(func() {
						Example("Originator Conversation ID")
					})
					Elem(func() {
						Example("19455-773836-1")
					})
				})
				Attribute("CreditPartyName", MapOf(String, String), func() {
					Key(func() {
						Example("CreditPartyName")
					})
					Elem(func() {
						Example("254708374149 - John Doe")
					})
				})
				Attribute("DebitPartyName", MapOf(String, String), func() {
					Key(func() {
						Example("DebitPartyName")
					})
					Elem(func() {
						Example("600134 - Safaricom157")
					})
				})
			})
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
