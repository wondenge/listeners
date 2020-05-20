// Code generated by goa v3.1.2, DO NOT EDIT.
//
// jenga service
//
// Command:
// $ goa gen github.com/wondenge/listeners/design

package jenga

import (
	"context"
)

// Service is the jenga service interface.
type Service interface {
	// Equity Bank Mpesa Async Callback
	Publish(context.Context, *JengaCallbackPayload) (res *JengaCallbackMedia, err error)
	// Account Alerts of monies in and out of your Equity Bank account
	Alerts(context.Context, *AccountAlerts) (res string, err error)
}

// ServiceName is the name of the service as defined in the design. This is the
// same value that is set in the endpoint request contexts under the ServiceKey
// key.
const ServiceName = "jenga"

// MethodNames lists the service method names as defined in the design. These
// are the same values that are set in the endpoint request contexts under the
// MethodKey key.
var MethodNames = [2]string{"publish", "alerts"}

// JengaCallbackPayload is the payload type of the jenga service publish method.
type JengaCallbackPayload struct {
	// Basic Authentication
	Auth *string
	// recipient phone number
	MobileNumber *string
	// recipient name
	Message *string
	// reference number
	Rrn *string
	// for M-Pesa the value is M-Pesa
	Service *string
	// M-Pesa receipt number
	TranID *string
	// M-Pesa return code
	ResultCode *string
	// M-Pesa return code description
	ResultDescription *string
	// Additional information
	AdditionalInfo *string
}

// JengaCallbackMedia is the result type of the jenga service publish method.
type JengaCallbackMedia struct {
	Code    *string
	Message *string
}

// AccountAlerts is the payload type of the jenga service alerts method.
type AccountAlerts struct {
	// Basic Authentication
	Auth         *string
	Name         *string
	MobileNumber *string
	// S2596405
	Reference *string
	// date
	Date *string
	// paymentMode
	PaymentMode *string
	// amount
	Amount *string
	Till   *string
	// billNumber
	BillNumber *string
	// orderAmount
	OrderAmount *string
	// serviceCharge
	ServiceCharge *string
	ServedBy      *string
	// additionalInfo
	AdditionalInfo *string
	// transactionType
	TransactionType *string
	// account
	Account *string
}
