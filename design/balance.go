package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

var AccountBalanceResult = Type("AccountBalanceResult", func() {
	Description("Account Balance Result")

	Extend(CommonResult)
	Attribute("MpesaResultParameters", func() {
		Extend(MpesaResultParameter)
	})
	Extend(ReferenceData)
})
