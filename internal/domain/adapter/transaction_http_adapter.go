package adapter

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/api"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain/service"
	log "github.com/sirupsen/logrus"
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
		log.Errorf("invalid payload %v", err)
		return err
	}
	nc := context.WithValue(ctx.Request().Context(), "external-auth", ctx.Request().Header["Authorization"][0])
	tid, err := tha.service.Confirmation(nc, ct)
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
