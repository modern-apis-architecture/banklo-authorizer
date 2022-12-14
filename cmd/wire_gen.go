// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/modern-apis-architecture/banklo-authorizer/internal/config"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain/adapter"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain/service"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/storage/http"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/storage/mongo"
)

// Injectors from wire.go:

func BuildAppContainer() (*adapter.TransactionHttpAdapter, error) {
	collection, err := mongo.ProvideCollection()
	if err != nil {
		return nil, err
	}
	mongoTransactionRepository := mongo.NewMongoTransactionRepository(collection)
	client := http.NewHttpClient()
	configConfig := config.ProvideConfig()
	httpExternalAuthorization := http.NewHttpExternalAuthorization(client, configConfig)
	externalAuthorization := service.NewExternalAuthorization(httpExternalAuthorization)
	cardService := service.NewCardService(client, configConfig)
	transactionService := service.NewTransactionService(mongoTransactionRepository, externalAuthorization, cardService)
	transactionHttpAdapter := adapter.NewTransactionHttpAdapter(transactionService)
	return transactionHttpAdapter, nil
}
