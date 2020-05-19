package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var B2CPaymentResult = Type("B2CPaymentResult", func() {

	// These parameters are always returned in the results whether the request was successful or not.
	Attribute("Result")

	// Global unique identifier for the transaction request returned
	// by the M-Pesa upon successful request submission.
	Attribute("ConversationId", String, func() {
		Description("Global unique identifier for the transaction request.")
		Example("236543-276372-2")
	})

	// Global unique identifier for the transaction request returned
	// by the API proxy upon successful request submission.
	Attribute("OriginatorConversationId", String, func() {
		Description("Global unique identifier for the transaction request.")
		Example("AG_2376487236_126732989KJHJKH")
	})

	// Message from the API that gives the status of the request processing
	// and usually maps to a specific result code value.
	Attribute("ResultDesc", String, func() {
		Description("Message from the API that gives the status of the request")
		Enum("Service request is has bee accepted successfully", "Initiator information is invalid")
	})

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

	// Unique M-PESA transaction ID for every payment request.
	// Same value is sent to customer over SMS upon successful processing.
	Attribute("TransactionID", String, func() {
		Description("Unique M-PESA transaction ID for every payment request.")
		Example("LHG31AA5TX")
	})




	// These parameters are only sent back if and only if the request was successful.

	Attribute("ResultParameters")
	Attribute("ResultParameter")

	// Same value is sent to customer over SMS upon successful processing.
	// It is usually returned under the ResultParameter array.
	Attribute("TransactionReceipt", String, func() {
		Description("This is a unique M-PESA transaction ID for every payment request.")
		Example("LHG31AA5TX")
	})

	// It is usually returned under the ResultParameter array.
	Attribute("TransactionAmount", Int, func() {
		Description("This is the amount that was transacted.")
		Example(100)
	})
	Attribute("B2CWorkingAccountAvailableFunds", Int64, func() {
		Description("Available balance of the Working account under the B2C shortcode used in the transaction.")
		Example(2000.0)
	})
	Attribute("B2CUtilityAccountAvailableFunds", Int64, func() {
		Description("Available balance of the Utility account under the B2C shortcode used in the transaction.")
		Example(23654.5)
	})
	Attribute("TransactionCompletedDateTime", String, func() {
		Description("This is the date and time that the transaction completed M-PESA.")
		Format(FormatDateTime)
		Example("01.08.2018 16:12:12")
	})
	Attribute("ReceiverPartyPublicName", String, func() {
		Description("This is the name and phone number of the customer who received the payment.")
		Example("254722000000 - Safaricom PLC")
	})
	Attribute("B2CChargesPaidAccountAvailableFunds", Int64, func() {
		Description("Available balance of the Charges Paid account under the B2C shortcode used in the transaction.")
		Example(236543.9)
	})
	Attribute("B2CRecipientIsRegisteredCustomer", String, func() {
		Description("This is a key that indicates whether the customer is a M-PESA registered customer or not")
		Enum(
			// "Y" for Yes
			"Y",
			//"N" for No
			"N",
		)
	})
})
