package auth

import (
	"context"
	"errors"

	"go-chat/config"
	"go-chat/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthRepo struct {
	users *mongo.Collection
}

func NewAuthRepo(db *mongo.Database) *AuthRepo {
	return &AuthRepo{users: db.Collection(config.C.DBKey.Users)}
}

func (r *AuthRepo) GetUserByUsername(username string, ctx context.Context) (models.User, error) {
	user := models.User{}
	err := r.users.FindOne(ctx, bson.M{"username": username}).Decode(&user)
	return user, err
}

func (r *AuthRepo) GetUserByID(id primitive.ObjectID, ctx context.Context) (models.User, error) {
	user := models.User{}
	err := r.users.FindOne(ctx, bson.M{"_id": id}).Decode(&user)
	return user, err
}

// CheckUserID checks if the user ID exists
func (r *AuthRepo) CheckUserID(id primitive.ObjectID, ctx context.Context) (bool, error) {
	count, err := r.users.CountDocuments(ctx, bson.M{"_id": id})
	return count > 0, err
}

func (r *AuthRepo) CheckUsername(username string, ctx context.Context) (bool, error) {
	count, err := r.users.CountDocuments(ctx, bson.M{"username": username})
	return count > 0, err
}

func (r *AuthRepo) CreateUser(user *models.User, ctx context.Context) (primitive.ObjectID, error) {
	result, err := r.users.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, err
	}

	id, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("failed to get inserted id")
	}
	return id, err
}
