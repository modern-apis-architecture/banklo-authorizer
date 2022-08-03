package adapter

import (
	"github.com/labstack/echo/v4"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/api"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain/service"
)

type TransactionHttpAdapter struct {
	service *service.TransactionService
}

func NewTransactionHttpAdapter(service *service.TransactionService) *TransactionHttpAdapter {
	return &TransactionHttpAdapter{service: service}
}

func (tha *TransactionHttpAdapter) CreateTransaction(ctx echo.Context) error {
	ct := &api.RequestTransaction{}
	if err := ctx.Bind(ct); err != nil {
		return err
	}
	tid, err := tha.service.Confirmation(ct)
	if err != nil {
		return err
	}
	return ctx.JSON(201, tid)
}

func (tha *TransactionHttpAdapter) RequestCancellation(ctx echo.Context, id string) error {
	return nil
}

func (tha *TransactionHttpAdapter) RequestReversal(ctx echo.Context, id string) error {
	return nil
}
