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

var ValidationAndConfirmation = Type("ValidationAndConfirmation", func() {
	Attribute("TransactionType", String, func() {
		Description("The transaction type specified during the payment request.")
		Enum("Buy Goods", "Pay Bill")
	})

	// This is the unique M-Pesa transaction ID for every payment request.
	// This is sent to both the call back messages and an confirmation SMS sent to customer.
	Attribute("TransID", String, func() {
		Description("Unique M-Pesa transaction ID for every payment request")
		Example("LHG31AA5TX")
	})

	// This is the Timestamp of the transaction, normaly in the formart
	// of YEAR+MONTH+DATE+HOUR+MINUTE+SECOND (YYYYMMDDHHMMSS)
	// Each part should be atleast two digits apart from the year which takes four digits.
	Attribute("TransTime", String, func() {
		Description("Timestamp of the transaction")
		Format(FormatDateTime)
		Example("20180713154301")
	})

	// This is the Amount transacted normaly a numeric value.
	// Money that customer pays to the Shorcode.
	// Only whole numbers are supported
	Attribute("TransAmount", Int, func() {
		Description("The Amount transacted normally a numeric value")
		Example(100)
	})

	// This is organizations shortcode (Paybill or Buygoods - A 5 to 6 digit account number)
	// used to identify an organization and receive the transaction.
	Attribute("BusinessShortCode", Int, func() {
		Description("Organizations shortcode (Paybill or Buygoods)")
		MinLength(5)
		MaxLength(6)
		Example(654321)
	})

	// 	This is the account number for which the customer is making the payment.
	//	 This is only applicable for Customer PayBill Transactions
	Attribute("BillRefNumber", String, func() {
		Description("The account number for which the customer is making the payment")
		MaxLength(20)
	})

	// The current utility account balance of the payment receiving organization shortcode.
	// For validation request, this field is usually blank whereas for the confirmation message,
	// the value represents the new balance after the payment has been received.
	Attribute("OrgAccountBalance", Int64, func() {
		Description("Current utility account balance of the payment receiving organization shortcode")
		Example(30671.00)
	})
	Attribute("MSISDN", Int, func() {
		Description("This is the mobile number of the customer making the payment.")
	})
	Attribute("FirstName", String, func() {
		Description("First Name of the customer making the payment, as per the M-Pesa register.")
		Example("John")
	})
	Attribute("MiddleName", String, func() {
		Description("Middle Name of the customer making the payment, as per the M-Pesa register.")
		Example("Doe")
	})
	Attribute("LastName", String, func() {
		Description("Last Name of the customer making the payment, as per the M-Pesa register.")
		Example("Jane")
	})
})



var LNMOCallbackResult = ResultType("LNMOCallbackResult", func() {
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
