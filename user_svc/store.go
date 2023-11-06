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
const dbCollection = "users"

type Store interface {
	Insert(ctx context.Context, user *types.User) (*types.User, error)
	Get(ctx context.Context, filter types.DBFilter) (*types.User, error)
	Update(ctx context.Context, filter types.DBFilter, params *UserDBUpdateParams) error
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]*types.User, error)
	Drop(context.Context) error
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

func (s *MongoStore) Drop(ctx context.Context) error {
	fmt.Println("--- dropping user collection")
	return s.coll.Drop(ctx)
}

func (s *MongoStore) Insert(ctx context.Context, user *types.User) (*types.User, error) {
	res, err := s.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}
	user.ID = res.InsertedID.(primitive.ObjectID).Hex()
	return user, nil
}

func (s *MongoStore) Get(ctx context.Context, filter types.DBFilter) (*types.User, error) {
	var user DBUser
	if filter["_id"] != nil {
		oid, err := primitive.ObjectIDFromHex(filter["_id"].(string))
		if err != nil {
			return nil, err
		}
		filter["_id"] = oid
	}
	if err := s.coll.FindOne(ctx, filter).Decode(&user); err != nil {
		return nil, err
	}
	return &types.User{
		ID: user.ID.Hex(),
		FirstName: user.FirstName,
		LastName: user.LastName,
		Email: user.Email,
		PhoneNumber: user.PhoneNumber,
		PhoneVerified: user.PhoneVerified,
		EncryptedPassword: user.EncryptedPassword,
		EmailVerified: user.EmailVerified,
	}, nil
}

func (s *MongoStore) Update(ctx context.Context, filter types.DBFilter, params *UserDBUpdateParams) error {
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

func (s *MongoStore) List(ctx context.Context) ([]*types.User, error) {
	cur, err := s.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	var users []*types.User
	if err = cur.All(ctx, &users); err != nil {
		return nil, err
	}
	return users, nil
}