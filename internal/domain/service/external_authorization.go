package service

import (
	"context"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain"
)

type ExternalAuthorization struct {
	extAuth ClientExternalAuthorization
}

func NewExternalAuthorization(extAuth ClientExternalAuthorization) *ExternalAuthorization {
	return &ExternalAuthorization{extAuth: extAuth}
}

func (ea *ExternalAuthorization) Authorize(ctx context.Context, t *domain.Transaction) (*domain.ExternalTransactionId, error) {
	nt, err := ea.extAuth.Authorize(ctx,t)
	if err != nil {
		return nil, err
	}
	return &domain.ExternalTransactionId{
		Id: nt.Id,
	}, err
}

func (ea *ExternalAuthorization) Cancellation(ctx context.Context,t *domain.Transaction) (*domain.ExternalTransactionId, error) {
	nt, err := ea.extAuth.Cancellation(ctx,t)
	if err != nil {
		return nil, err
	}
	return &domain.ExternalTransactionId{Id: nt.Id}, err

}

func (ea *ExternalAuthorization) Reversal(ctx context.Context,t *domain.Transaction) (*domain.ExternalTransactionId, error) {
	nt, err := ea.extAuth.Reversal(ctx,t)
	if err != nil {
		return nil, err
	}
	return &domain.ExternalTransactionId{Id: nt.Id}, err
}
