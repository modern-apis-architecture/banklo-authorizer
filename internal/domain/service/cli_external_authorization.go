package service

import "github.com/modern-apis-architecture/banklo-authorizer/internal/domain"

type ClientExternalAuthorization interface {
	Authorize(t *domain.Transaction) (*domain.ExternalTransactionResponse, error)
	Cancellation(t *domain.Transaction) (*domain.ExternalTransactionResponse, error)
	Reversal(t *domain.Transaction) (*domain.ExternalTransactionResponse, error)
}
