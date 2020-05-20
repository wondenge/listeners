package listener

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	jenga "github.com/wondenge/listeners/gen/jenga"
)

// jenga service example implementation.
// The example methods log the requests and return zero values.
type jengasrvc struct {
	logger log.Logger
}

// NewJenga returns the jenga service implementation.
func NewJenga(logger log.Logger) jenga.Service {
	return &jengasrvc{logger}
}

// Equity Bank Mpesa Async Callback
func (s *jengasrvc) Publish(ctx context.Context, p *jenga.JengaCallbackPayload) (res *jenga.JengaCallbackMedia, err error) {
	res = &jenga.JengaCallbackMedia{}
	s.logger.Log("info", fmt.Sprintf("jenga.publish"))
	return
}

// Account Alerts of monies in and out of your Equity Bank account
func (s *jengasrvc) Alerts(ctx context.Context, p *jenga.AccountAlerts) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("jenga.alerts"))
	return
}
