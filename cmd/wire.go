//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/config"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain/adapter"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain/repository"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain/service"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/storage/http"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/storage/mongo"
)

func BuildAppContainer() (*adapter.TransactionHttpAdapter, error) {
	wire.Build(mongo.ProvideCollection, mongo.NewMongoTransactionRepository, http.NewHttpClient, http.NewHttpExternalAuthorization,
		config.ProvideConfig, service.NewCardService,
		wire.Bind(new(repository.TransactionRepository), new(*mongo.MongoTransactionRepository)),
		wire.Bind(new(service.ClientExternalAuthorization), new(*http.HttpExternalAuthorization)),
		service.NewExternalAuthorization, service.NewTransactionService, adapter.NewTransactionHttpAdapter,
	)
	return nil, nil
}
