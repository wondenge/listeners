package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var TransactionStatusResult = Type("TransactionStatusResult", func() {
	Description("Transaction Status Result")

	Extend(CommonResult)
	Attribute("ResultParameters", func() {
		Extend(MpesaResultParameter)
	})
	Extend(ReferenceData)
})
