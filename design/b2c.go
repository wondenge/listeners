package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var B2CPaymentResult = Type("B2CPaymentResult", func() {
	Description("Result Parameters")

	// "Result" object is always returned in the results
	// whether the request was successful or not.
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
	})

	// "ResultParameters" object is only sent back if and only
	// if the request was successful.
	// It's a JSON object holding more details for the transaction.
	Attribute("ResultParameters", func() {
		Attribute("ResultParameter", ArrayOf(B2CResultParameters))
	})
	Attribute("ReferenceData", func() {
		Attribute("ReferenceItem", MapOf(String, String), func() {
			Key(func() {
				MinLength(1)
				Example("QueueTimeoutURL")
			})
			Elem(func() {
				Pattern("[a-zA-Z]+")
				Example("https://internalsandbox.safaricom.co.ke/mpesa/b2cresults/v1/submit")
			})
		})
	})
})

var B2CResultParameters = Type("B2CResultParameters", func() {

	// Same value is sent to customer over SMS upon successful processing.
	// It is usually returned under the ResultParameter array.
	Attribute("TransactionReceipt", String, func() {
		Description("This is a unique M-PESA transaction ID for every payment request.")
		Pattern("[a-zA-Z]+")
		Example("LGH3197RIB")
	})
	// It is usually returned under the ResultParameter array.
	Attribute("TransactionAmount", Int, func() {
		Description("This is the amount that was transacted.")
		Example(8000) // Number

	})
	Attribute("B2CWorkingAccountAvailableFunds", Float64, func() {
		Description("Available balance of the Working account under the B2C shortcode used in the transaction.")
		Example(2000.0) // Decimal
	})
	Attribute("B2CUtilityAccountAvailableFunds", Float64, func() {
		Description("Available balance of the Utility account under the B2C shortcode used in the transaction.")
		Example(23654.5) // Decimal
	})
	Attribute("TransactionCompletedDateTime", String, func() {
		Description("This is the date and time that the transaction completed M-PESA.")
		Format(FormatDateTime)
		Example("17.07.2017 10:54:57")
	})
	Attribute("ReceiverPartyPublicName", String, func() {
		Description("This is the name and phone number of the customer who received the payment.")
		Pattern("[a-zA-Z]+")
		Example("254722000000 - Safaricom PLC")
	})
	Attribute("B2CChargesPaidAccountAvailableFunds", Float64, func() {
		Description("Available balance of the Charges Paid account under the B2C shortcode used in the transaction.")
		Example(236543.9) // Decimal
	})
	Attribute("B2CRecipientIsRegisteredCustomer", String, func() {
		Description("Key indicating whether customer is registered on M-PESA")
		Pattern("[a-zA-Z]+")
		// "Y" for Yes
		//"N" for No
		Enum("Y", "N")
		Example("Y")
	})
})
