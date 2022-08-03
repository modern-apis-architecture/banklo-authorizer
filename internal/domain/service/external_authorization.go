package service

import "github.com/modern-apis-architecture/banklo-authorizer/internal/domain"

type ExternalAuthorization struct {
	extAuth ClientExternalAuthorization
}

func NewExternalAuthorization(extAuth ClientExternalAuthorization) *ExternalAuthorization {
	return &ExternalAuthorization{extAuth: extAuth}
}

func (ea *ExternalAuthorization) Authorize(t *domain.Transaction) {

}
