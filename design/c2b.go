package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var ValidationResult = Type("ValidationResult", func() {
	Attribute("TransactionType", String, func() {
		Description("Transaction type specified during the payment request.")
		Enum("Buy Goods", "Pay Bill")
	})

	// This is sent to both the call back messages
	// and an confirmation SMS sent to customer.
	Attribute("TransID", String, func() {
		Description("Unique M-Pesa transaction ID for every payment request")
		Example("LHG31AA5TX")
	})

	// YEAR+MONTH+DATE+HOUR+MINUTE+SECOND (YYYYMMDDHHMMSS)
	// Each part should be at least two digits apart from
	// the year which takes four digits.
	Attribute("TransTime", String, func() {
		Description("Timestamp of the transaction")
		Format(FormatDateTime)
		Example("20180713154301")
	})

	// Money that customer pays to the Shortcode.
	// Only whole numbers are supported
	Attribute("TransAmount", Int, func() {
		Description("The Amount transacted normally a numeric value")
		Example(100)
	})

	Attribute("BusinessShortCode", Int, func() {
		Description("Organizations shortcode (Paybill or Buygoods)")
		Example(654321)
	})

	//	This is only applicable for Customer PayBill Transactions
	Attribute("BillRefNumber", String, func() {
		Description("Account number for which the customer is making the payment")
		MaxLength(20)
	})
	Attribute("InvoiceNumber", String)

	// For validation request, this field is usually blank
	// whereas for the confirmation message, the value represents
	// the new balance after the payment has been received.
	Attribute("OrgAccountBalance", Int, func() {
		Description("Current utility account balance of the payment receiving organization shortcode")
		Example(30671)
	})
	Attribute("ThirdPartyTransID", String)
	Attribute("MSISDN", Int, func() {
		Description("Mobile number of the customer making the payment.")
	})
	Attribute("FirstName", String, func() {
		Description("First Name of the customer making the payment.")
		Example("John")
	})
	Attribute("MiddleName", String, func() {
		Description("Middle Name of the customer making the payment.")
		Example("Doe")
	})
	Attribute("LastName", String, func() {
		Description("Last Name of the customer making payment.")
		Example("Jane")
	})
})

var ConfirmationResult = Type("ConfirmationResult", func() {
	Attribute("TransactionType", String, func() {
		Description("Transaction type specified during the payment request.")
		Enum("Buy Goods", "Pay Bill")
	})

	// This is sent to both the call back messages
	// and an confirmation SMS sent to customer.
	Attribute("TransID", String, func() {
		Description("Unique M-Pesa transaction ID for every payment request")
		Example("LHG31AA5TX")
	})

	// YEAR+MONTH+DATE+HOUR+MINUTE+SECOND (YYYYMMDDHHMMSS)
	// Each part should be at least two digits apart from
	// the year which takes four digits.
	Attribute("TransTime", String, func() {
		Description("Timestamp of the transaction")
		Format(FormatDateTime)
		Example("20180713154301")
	})

	// Money that customer pays to the Shortcode.
	// Only whole numbers are supported
	Attribute("TransAmount", Int, func() {
		Description("The Amount transacted normally a numeric value")
		Example(100)
	})

	Attribute("BusinessShortCode", Int, func() {
		Description("Organizations shortcode (Paybill or Buygoods)")
		Example(654321)
	})

	//	This is only applicable for Customer PayBill Transactions
	Attribute("BillRefNumber", String, func() {
		Description("Account number for which the customer is making the payment")
		MaxLength(20)
	})
	Attribute("InvoiceNumber", String)

	// For validation request, this field is usually blank
	// whereas for the confirmation message, the value represents
	// the new balance after the payment has been received.
	Attribute("OrgAccountBalance", Int, func() {
		Description("Current utility account balance of the payment receiving organization shortcode")
		Example(30671)
	})
	Attribute("ThirdPartyTransID", String)
	Attribute("MSISDN", Int, func() {
		Description("Mobile number of the customer making the payment.")
	})
	Attribute("FirstName", String, func() {
		Description("First Name of the customer making the payment.")
		Example("John")
	})
	Attribute("MiddleName", String, func() {
		Description("Middle Name of the customer making the payment.")
		Example("Doe")
	})
	Attribute("LastName", String, func() {
		Description("Last Name of the customer making payment.")
		Example("Jane")
	})
})
