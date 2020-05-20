package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var CommonResult = Type("CommonResult", func() {
	Description("Common Result for Endpoints")

	// "Result" object is always returned in the results
	// whether the request was successful or not.
	Extend(MpesaResult)
})

var MpesaResult = Type("MpesaResult", func() {
	Attribute("MpesaResultType", Int, func() {
		Description("Status code indicating whether transaction was already sent to your listener")
		Example(0)
	})

	// 0 means success.
	// Any other code means an error occurred
	// or the transaction failed.
	Attribute("MpesaResultCode", Int, func() {
		Description("Numeric status code indicating the status of the transaction processing")
		Example(0)
	})

	// Usually maps to a specific result code value
	Attribute("MpesaResultDesc", String, func() {
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

var MpesaResultParameter = Type("MpesaResultParameter", func() {

	// ACCOUNT BALANCE
	Attribute("AccountBalance", String, func() {
		Description("Account Balance")
		Example("Working Account|KES|46713.00|46713.00|0.00|0.00&Float Account|KES|0.00|0.00|0.00|0.00&Utility Account|KES|49217.00|49217.00|0.00|0.00&Charges Paid Account|KES|-220.00|-220.00|0.00|0.00&Organization Settlement Account|KES|0.00|0.00|0.00|0.00")
	})
	Attribute("BOCompletedTime", String, func() {
		Description("BO Completed Time")
		Format(FormatDateTime)
		Example("20170728095642")
	})

	// TRANSACTION STATUS
	Attribute("ReceiptNo", String, func() {
		Description("Receipt No")
		Example("LGR919G2AV")
	})
	Attribute("ConversationID", String, func() {
		Description("Conversation ID")
		Example("AG_20170727_00004492b1b6d0078fbe")
	})
	Attribute("FinalisedTime", String, func() {
		Description("Finalised Time")
		Format(FormatDateTime)
		Example("20170727101415")
	})
	Attribute("Amount", Int, func() {
		Description("Amount")
		Example(10)
	})
	Attribute("TransactionStatus", String, func() {
		Description("Transaction Status")
		Example("Completed")
	})
	Attribute("ReasonType", String, func() {
		Description("Reason Type")
		Example("Salary Payment via API")
	})
	Attribute("TransactionReason", String, func() {
		Example("Transaction Reason")
	})
	Attribute("DebitPartyCharges", String, func() {
		Description("Debit Party Charges")
		Example("Fee For B2C Payment|KES|33.00")
	})
	Attribute("DebitAccountType", String, func() {
		Description("Debit Account Type")
		Example("Utility Account")
	})
	Attribute("InitiatedTime", String, func() {
		Description("Initiated Time")
		Format(FormatDateTime)
		Example("20170727101415")
	})
	Attribute("OriginatorConversationID", String, func() {
		Description("Originator Conversation ID")
		Example("19455-773836-1")
	})
	Attribute("CreditPartyName", String, func() {
		Description("Credit Party Name")
		Example("254708374149 - John Doe")
	})
	Attribute("DebitPartyName", String, func() {
		Description("Debit Party Name")
		Example("600134 - Safaricom157")
	})

	// C2B

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

var ReferenceData = Type("ReferenceData", func() {
	Extend(ReferenceItem)
})

var ReferenceItem = Type("ReferenceItem", func() {
	Attribute("QueueTimeoutURL", String, func() {
		Description("Queue Timeout URL")
		Example("https://internalsandbox.safaricom.co.ke/mpesa/abresults/v1/submit")
	})
	Attribute("Occasion", func() {
		Example("Occasion")
	})
})
