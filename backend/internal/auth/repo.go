package auth

import (
	"context"
	"go-chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type authRepo struct {
	db *mongo.Collection
}

func NewAuthRepo(db *mongo.Database) *authRepo {
	return &authRepo{db: db.Collection("users")}
}

func (r *authRepo) GetUserByUsername(username string, ctx context.Context) (models.User, error) {
	user := models.User{}
	err := r.db.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user, err
}

func (r *authRepo) CheckUsername(username string, ctx context.Context) (bool, error) {
	count, err := r.db.CountDocuments(ctx, bson.M{"username": username})
	return count > 0, err
}

func (r *authRepo) CreateUser(user *models.User, ctx context.Context) (primitive.ObjectID, error) {
	result, err := r.db.InsertOne(ctx, user)
	return result.InsertedID.(primitive.ObjectID), err
}
