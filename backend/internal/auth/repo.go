package auth

import (
	"context"
	"go-chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (r *AuthRepo) CheckUsername(username string, ctx context.Context) (bool, error) {
	count, err := r.users.CountDocuments(ctx, bson.M{"username": username})
	return count > 0, err
}

func (r *AuthRepo) CreateUser(user *models.User, ctx context.Context) (primitive.ObjectID, error) {
	result, err := r.users.InsertOne(ctx, user)
	return result.InsertedID.(primitive.ObjectID), err
}
