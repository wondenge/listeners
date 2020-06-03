package design

import (
	. "goa.design/goa/v3/dsl"
	_ "goa.design/plugins/v3/docs"      // Generates documentation
	_ "goa.design/plugins/v3/goakit"    // Enables goakit
	_ "goa.design/plugins/v3/zaplogger" // Enables ZapLogger Plugin
)

// API describes the global properties of the API server.
var _ = API("safaricom", func() {

	// API title.
	Title("M-pesa HTTP Listeners")

	// Description of API
	Description("HTTP service for interacting with callbacks supporting M-pesa APIs")

	// Version of API
	Version("1.0")

	// Terms of use of API
	TermsOfService("https://github.com/wondenge/mpesa-notifiers/blob/master/LICENSE")

	// Contact info for Author and Maintainer
	Contact(func() {
		Name("William Ondenge")
		Email("ondengew@gmail.com")
		URL("https://www.ondenge.me")
	})

	// License for Library usage
	License(func() {
		Name("Apache License")
		URL("https://github.com/wondenge/mpesa-notifiers/blob/master/LICENSE")
	})

	// Library Documentation
	Docs(func() {
		Description("Library Usage")
		URL("https://github.com/wondenge/mpesa-notifiers/blob/master/README.md")
	})

	// Server describes a single process listening for client requests.
	Server("daraja", func() {
		Description("daraja hosts all callback services supporting M-pesa APIs.")

		// List services hosted by this server.
		Services("mpesa", "health", "swagger")

		// List the Hosts and their transport URLs.
		Host("local", func() {
			Description("local host")
			URI("http://localhost:3000")
		})

		Host("docker", func() {
			Description("docker host")

			// We use 0.0.0.0 so the bind will work in the docker image.
			// We will not be doing TLS here as we expect the instance to
			// be running in a firewalled environment behind a load balancer
			// where the load balancer kubernetes ingress will be responsible for TLS.
			URI("http://0.0.0.0:8000")
		})
	})
})

var _ = Service("health", func() {

	// HTTP defines the HTTP transport specific
	// properties of an API, a service or a single method.
	HTTP(func() {
		Path("/health")
	})

	// Defines a single service method
	Method("show", func() {
		Description("Health check endpoint.")
		Result(String)
		HTTP(func() {
			GET("/")
			Response(func() {
				Code(StatusOK)
				ContentType("text/plain")
			})
		})
	})
})

var _ = Service("swagger", func() {
	Description("The swagger service serves the API swagger definition.")
	HTTP(func() {
		Path("/swagger")
	})

	// Serve the file with relative path
	// ../../gen/http/openapi.json for requests sent to /swagger.json.
	Files("/swagger.json", "../../gen/http/openapi.json", func() {
		Description("JSON document containing the API swagger definition")
	})
})
