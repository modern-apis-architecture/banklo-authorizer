package http

import (
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain"
	"net/http"
)

type HttpExternalAuthorization struct {
	cli *http.Client
}

func NewHttpExternalAuthorization(cli *http.Client) *HttpExternalAuthorization {
	return &HttpExternalAuthorization{cli: cli}
}

func (hea *HttpExternalAuthorization) Authorize(t *domain.Transaction) error {
	return nil
}

func (hea *HttpExternalAuthorization) Cancellation(t *domain.Transaction) error {
	return nil
}

func (hea *HttpExternalAuthorization) Reversal(t *domain.Transaction) error {
	return nil
}
