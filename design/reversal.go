package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// Note the TransactionID. This is the transaction ID of the reversal
// request itself, not the original transaction which was being reversed.
// The reversal request itself gets its own transaction ID.
var ReversalResult = Type("ReversalResult", func() {
	Description("Successful Reversal Callback")

	Extend(CommonResult)
	Extend(ReferenceData)
})
