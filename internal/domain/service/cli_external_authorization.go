package service

import "github.com/modern-apis-architecture/banklo-authorizer/internal/domain"

type ClientExternalAuthorization interface {
	Authorize(t *domain.Transaction) error
	Cancellation(t *domain.Transaction) error
	Reversal(t *domain.Transaction) error
}