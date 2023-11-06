package main

import (
	"Jimbo8702/randomThoughts/diggity-dawg/types"
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"gopkg.in/mgo.v2/bson"
)

const MongoDBNameEnvName = "MONGO_DB_NAME"
const dbCollection = "accounts"

type Store interface {
	Insert(ctx context.Context, acc *types.Account) (*types.Account, error)
	Get(ctx context.Context, filter types.DBFilter) (*types.Account, error)
	Update(ctx context.Context, filter types.DBFilter, params *AccountDBUpdateParams) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*types.Account, error)
	Drop(ctx context.Context) error
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

func (s *MongoStore) Insert(ctx context.Context, acc *types.Account) (*types.Account, error) {
	res, err := s.coll.InsertOne(ctx, acc)
	if err != nil {
		return nil, err
	}
	acc.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return acc, nil
}

func (s *MongoStore) Get(ctx context.Context, filter types.DBFilter) (*types.Account, error) {
	var acc DBAccount
	if filter["_id"] != nil {
		oid, err := primitive.ObjectIDFromHex(filter["_id"].(string))
		if err != nil {
			return nil, err
		}
		filter["_id"] = oid
	}
	if err := s.coll.FindOne(ctx, filter).Decode(&acc); err != nil {
		return nil, err
	}
	return &types.Account{
		ID: acc.ID.Hex(),
		UserID: acc.UserID.Hex(),
		StripeCustomerID: acc.StripeCustomerID,
		SubscriptionStatus: acc.SubscriptionStatus,
		StripeSubscriptionID: acc.StripeSubscriptionID,
	}, nil
}

func (s *MongoStore) Update(ctx context.Context, filter types.DBFilter, params *AccountDBUpdateParams) error {
	if filter["_id"] != nil {
		oid, err := primitive.ObjectIDFromHex(filter["_id"].(string))
		if err != nil {
			return err
		}
		filter["_id"] = oid
	}
	update := bson.M{"$set": params.ToBSON()}
	_, err := s.coll.UpdateOne(ctx, filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoStore) Delete(ctx context.Context, id string) error {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = s.coll.DeleteOne(ctx, bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}

func (s *MongoStore) List(ctx context.Context) ([]*types.Account, error) {
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var accounts []*types.Account
	if err = cur.All(ctx, &accounts); err != nil {
		return nil, err
	}
	return accounts, nil
}
func (s *MongoStore) Drop(ctx context.Context) error { 
	fmt.Println("--- dropping user collection")
	return s.coll.Drop(ctx)
}