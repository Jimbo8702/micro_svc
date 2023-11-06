package main

import (
	"Jimbo8702/randomThoughts/diggity-dawg/types"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

const MongoDBNameEnvName = "MONGO_DB_NAME"
const dbCollection = "users"

type Dropper interface {
	Drop(context.Context) error
}

type Store interface {
	Dropper

	Insert(ctx context.Context, user *types.User) (*types.User, error)
	Get(ctx context.Context, filter DBFilter) (*types.User, error)
	Update(ctx context.Context, filter DBFilter, params UserDBUpdateParams) error
	Delete(ctx context.Context, id string) error

	List(ctx context.Context) ([]*types.User, error)
	GetById(ctx context.Context, id string) (*types.User, error)
}

type MongoStore struct {
	client 	*mongo.Client
	coll  	*mongo.Collection
}

func NewMongoStore(client *mongo.Client) *MongoStore {
	dbName := os.Getenv(MongoDBNameEnvName)
	return &MongoStore{
		client: client,
		coll: client.Database(dbName).Collection(dbCollection),
	}
}

func (s *MongoStore) Insert(ctx context.Context, user *types.User) (*types.User, error)
func (s *MongoStore) Get(ctx context.Context, filter DBFilter) (*types.User, error)
func (s *MongoStore) Update(ctx context.Context, filter DBFilter, params UserDBUpdateParams) error
func (s *MongoStore) Delete(ctx context.Context, id string) error
func (s *MongoStore) List(ctx context.Context) ([]*types.User, error)
func (s *MongoStore) GetById(ctx context.Context, id string) (*types.User, error)