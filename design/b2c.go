package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var B2CPaymentResult = Type("B2CPaymentResult", func() {
	Description("Result Parameters")

	Extend(CommonResult)

	// "ResultParameters" object is only sent back if and only
	// if the request was successful.
	// It's a JSON object holding more details for the transaction.
	Attribute("MpesaResultParameters", func() {
		Extend(MpesaResultParameter)
	})
	Extend(ReferenceData)
})
