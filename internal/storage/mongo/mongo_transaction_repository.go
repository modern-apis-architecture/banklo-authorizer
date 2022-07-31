package mongo

import (
	"context"
	"fmt"
	"github.com/modern-apis-architecture/banklo-authorizer/internal/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoTransactionRepository struct {
	collection *mongo.Collection
}

func (mtr *MongoTransactionRepository) Store(transaction *domain.Transaction) (*domain.TransactionId, error) {
	ctx := context.Background()
	opts := options.InsertOne()
	doc, err := mtr.collection.InsertOne(ctx, transaction, opts)
	if err != nil {
		return nil, fmt.Errorf("could not save document to mongo: %w", err)
	}
	tid := fmt.Sprintf("%v", doc.InsertedID)
	return domain.NewTransactionId(tid), nil
}

func NewMongoTransactionRepository(collection *mongo.Collection) *MongoTransactionRepository {
	return &MongoTransactionRepository{collection: collection}
}
