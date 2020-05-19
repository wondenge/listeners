package listener

import (
	"context"
	"fmt"

	"github.com/go-kit/kit/log"
	mpesa "github.com/wondenge/listeners/gen/mpesa"
)

// mpesa service example implementation.
// The example methods log the requests and return zero values.
type mpesasrvc struct {
	logger log.Logger
}

// NewMpesa returns the mpesa service implementation.
func NewMpesa(logger log.Logger) mpesa.Service {
	return &mpesasrvc{logger}
}

// Account Balance Queue TimeOut URL
func (s *mpesasrvc) AccountBalanceTimeout(ctx context.Context, p *mpesa.AccountBalanceResult) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("mpesa.AccountBalanceTimeout"))
	return
}

// Account Balance Result URL
func (s *mpesasrvc) AccountBalanceResultEndpoint(ctx context.Context, p *mpesa.AccountBalanceResult) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("mpesa.AccountBalanceResult"))
	return
}

// Transaction Status Queue TimeOut URL
func (s *mpesasrvc) TransactionStatusTimeout(ctx context.Context, p *mpesa.TransactionStatusResult) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("mpesa.TransactionStatusTimeout"))
	return
}

// Transaction Status Result URL
func (s *mpesasrvc) TransactionStatusResultEndpoint(ctx context.Context, p *mpesa.TransactionStatusResult) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("mpesa.TransactionStatusResult"))
	return
}

// Reversal Queue TimeOut URL
func (s *mpesasrvc) ReversalTimeout(ctx context.Context, p *mpesa.ReversalResult) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("mpesa.ReversalTimeout"))
	return
}

// Reversal Result URL
func (s *mpesasrvc) ReversalResultEndpoint(ctx context.Context, p *mpesa.ReversalResult) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("mpesa.ReversalResult"))
	return
}

// B2C Queue TimeOut URL
func (s *mpesasrvc) B2CTimeout(ctx context.Context, p *mpesa.B2CPaymentResult) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("mpesa.B2CTimeout"))
	return
}

// B2C Result URL
func (s *mpesasrvc) B2CResult(ctx context.Context, p *mpesa.B2CPaymentResult) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("mpesa.B2CResult"))
	return
}

// C2B Validation URL
func (s *mpesasrvc) C2BValidation(ctx context.Context, p *mpesa.ValidationResult) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("mpesa.C2BValidation"))
	return
}

// C2B Confirmation URL
func (s *mpesasrvc) C2BConfirmation(ctx context.Context, p *mpesa.ConfirmationResult) (res string, err error) {
	s.logger.Log("info", fmt.Sprintf("mpesa.C2BConfirmation"))
	return
}
