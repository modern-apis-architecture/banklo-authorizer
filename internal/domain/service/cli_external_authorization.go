package service

import (
	"context"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain"
)

type ClientExternalAuthorization interface {
	Authorize(ctx context.Context, t *domain.Transaction) (*domain.ExternalTransactionResponse, error)
	Cancellation(ctx context.Context, t *domain.Transaction) (*domain.ExternalTransactionResponse, error)
	Reversal(ctx context.Context, t *domain.Transaction) (*domain.ExternalTransactionResponse, error)
}
