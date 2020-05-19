package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var B2CPaymentResult = Type("B2CPaymentResult", func() {

	// These parameters are always returned in the results whether the request was successful or not.
	Attribute("Result", func() {
		// Status code that indicates whether the transaction was already sent
		// to your listener. Usual value is 0.
		Attribute("ResultType", Int, func() {
			Description("Status code that indicates whether the transaction was already sent")
			Default(0)
		})
		// This is a numeric status code that indicates the status of the transaction processing.
		// 0 means success and any other code means an error occured or the transaction failed.
		Attribute("ResultCode", Int, func() {
			Description("Numeric status code that indicates the status of the transaction processing")
			Example(0)
		})
		// Message from the API that gives the status of the request processing
		// and usually maps to a specific result code value.
		Attribute("ResultDesc", String, func() {
			Description("Message from the API that gives the status of the request")
			Enum("Service request is has bee accepted successfully", "Initiator information is invalid")
		})
		// Global unique identifier for the transaction request returned
		// by the API proxy upon successful request submission.
		Attribute("OriginatorConversationId", String, func() {
			Description("Global unique identifier for the transaction request.")
			Example("AG_2376487236_126732989KJHJKH")
		})
		// Global unique identifier for the transaction request returned
		// by the M-Pesa upon successful request submission.
		Attribute("ConversationId", String, func() {
			Description("Global unique identifier for the transaction request.")
			Example("236543-276372-2")
		})
		// Unique M-PESA transaction ID for every payment request.
		// Same value is sent to customer over SMS upon successful processing.
		Attribute("TransactionID", String, func() {
			Description("Unique M-PESA transaction ID for every payment request.")
			Example("LHG31AA5TX")
		})

	})
	Attribute("ResultParameters", func() {
		Attribute("ResultParameter", func() {

			// Same value is sent to customer over SMS
			// upon successful processing. It is usually
			// returned under the ResultParameter array.
			Attribute("TransactionReceipt", MapOf(String, String), func() {
				Description("This is a unique M-PESA transaction ID for every payment request.")
				Key(func() {
					MinLength(1)
					Example("TransactionReceipt")
				})
				Elem(func() {
					Pattern("[a-zA-Z]+")
					Example("LGH3197RIB")
				})
			})
			// It is usually returned under the ResultParameter array.
			Attribute("TransactionAmount", MapOf(String, Int), func() {
				Description("This is the amount that was transacted.")
				Key(func() {
					MinLength(1)
					Example("TransactionAmount")
				})
				Elem(func() {
					Example(8000)
				})
			})
			Attribute("B2CWorkingAccountAvailableFunds", MapOf(String, Int), func() {
				Description("Available balance of the Working account under the B2C shortcode used in the transaction.")
				Key(func() {
					MinLength(1)
					Example("B2CWorkingAccountAvailableFunds")
				})
				Elem(func() {
					Example(150000)
				})
			})
			Attribute("B2CUtilityAccountAvailableFunds", MapOf(String, Int), func() {
				Description("Available balance of the Utility account under the B2C shortcode used in the transaction.")
				Key(func() {
					MinLength(1)
					Example("B2CUtilityAccountAvailableFunds")
				})
				Elem(func() {
					Example(133568)
				})
			})
			Attribute("TransactionCompletedDateTime", MapOf(String, String), func() {
				Description("This is the date and time that the transaction completed M-PESA.")
				Key(func() {
					MinLength(1)
					Example("TransactionCompletedDateTime")
				})
				Elem(func() {
					Format(FormatDateTime)
					Example("17.07.2017 10:54:57")
				})
			})
			Attribute("ReceiverPartyPublicName", MapOf(String, String), func() {
				Description("This is the name and phone number of the customer who received the payment.")
				Key(func() {
					MinLength(1)
					Example("ReceiverPartyPublicName")
				})
				Elem(func() {
					Pattern("[a-zA-Z]+")
					Example("254708374149 - John Doe")
				})
			})
			Attribute("B2CChargesPaidAccountAvailableFunds", MapOf(String, Int), func() {
				Description("Available balance of the Charges Paid account under the B2C shortcode used in the transaction.")
				Key(func() {
					MinLength(1)
					Example("B2CChargesPaidAccountAvailableFunds")
				})
				Elem(func() {
					Minimum(0)
					Example(0)
				})
			})
			Attribute("B2CRecipientIsRegisteredCustomer", MapOf(String, String), func() {
				Description("This is a key that indicates whether the customer is a M-PESA registered customer or not")
				Key(func() {
					MinLength(1)
					Example("ReceiverPartyPublicName")
				})
				Elem(func() {
					Pattern("[a-zA-Z]+")
					// "Y" for Yes
					//"N" for No
					Enum("Y", "N")
					Example("Y")
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
				Example("https://internalsandbox.safaricom.co.ke/mpesa/b2cresults/v1/submit")
			})
		})
	})
})
