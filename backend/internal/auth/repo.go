package auth

import (
	"context"
	"errors"
	"go-chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type AuthRepo struct {
	users *mongo.Collection
}

func NewAuthRepo(db *mongo.Database) *AuthRepo {
	return &AuthRepo{users: db.Collection("users")}
}

func (r *AuthRepo) GetUserByUsername(username string, ctx context.Context) (models.User, error) {
	user := models.User{}
	result := r.users.FindOne(ctx, bson.M{"username": username})
	if result.Err() != nil {
		return user, result.Err()
	}
	err := result.Decode(&user)
	return user, err
}

// CheckUsername checks if the username is exist
func (r *AuthRepo) CheckUsername(username string, ctx context.Context) (bool, error) {
	count, err := r.users.CountDocuments(ctx, bson.M{"username": username})
	return count > 0, err
}

func (r *AuthRepo) CreateUser(user *models.User, ctx context.Context) (primitive.ObjectID, error) {
	result, err := r.users.InsertOne(ctx, user, &options.InsertOneOptions{})
	if err != nil {
		return primitive.NilObjectID, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("failed to get inserted id")
	}
	return id, err
}
