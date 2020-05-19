package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("listener", func() {
	Title("Listener")
	Description("Golang HTTP service listening for Callback and Timeout URLs")
	Version("1.0")
	TermsOfService("https://github.com/wondenge/mpesa-listener/blob/master/LICENSE")
	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
		URL("https://www.ondenge.me")
	})
	License(func() {
		Name("Apache License")
		URL("https://github.com/wondenge/mpesa-listener/blob/master/LICENSE")
	})
	Docs(func() {
		Description("Library Usage")
		URL("https://github.com/wondenge/mpesa-listener/blob/master/README.md")
	})
	Server("listener", func() {
		Description("listener hosts all Listener Services.")
		Services("mpesa")
		Host("local", func() {
			Description("Localhost")
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("mpesa", func() {

})

// '/mpesa/b2c/v1'
// '/mpesa/b2b/v1'

// B2C CALLBACK '/b2c/result'
// B2C ResultURL - /api/v1/b2c/result

// B2C TIMEOUT '/b2c/timeout'
// B2C QueueTimeoutURL - /api/v1/b2c/timeout

// C2B VALIDATION REQUEST '/validation'
// C2B ValidationURL - /api/v1/c2b/validation

// C2B CONFIRMATION REQUEST- '/confirmation'
// C2B ConfirmationURL - /api/v1/c2b/confirmation

// B2B CALLBACK '/b2b/result'
// B2B ResultURL - /api/v1/b2b/result

// B2B TIMEOUT '/b2b/timeout'
// B2B QueueTimeoutURL - /api/v1/b2b/timeout
