package user

import (
	"go-chat/config"

	"go.mongodb.org/mongo-driver/mongo"
)

type userRepo struct {
	users, messages *mongo.Collection
}

func NewUserRepo(db *mongo.Database) *userRepo {
	return &userRepo{
		users:    db.Collection(config.C.DBKey.Users),
		messages: db.Collection(config.C.DBKey.Messages),
	}
}

// TODO research the use of generic methods to write/read user and group messages
// func (r *userRepo) GetReleatedUsers(userID primitive.ObjectID, ctx context.Context) ([]models.User, error) {
// 	messages := []models.UserMessage{}
// 	groupMessages := []models.GroupMessage{}
// 	cursor1, err := r.messages.Find(ctx, bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	cursor2, err := r.groupMessages.Find(ctx, bson.D{})
// 	if err != nil {
// 		return nil, err
// 	}
// 	if err = cursor1.All(ctx, &messages); err != nil {
// 		return nil, err
// 	}
// 	if err = cursor2.All(ctx, &groupMessages); err != nil {
// 		return nil, err
// 	}
// 	users := []models.User{}
// 	for _, message := range messages {
// 		if message.SenderID != userID {

// 		}
// 		if message.To != userID {
// 			users[message.To] = models.User{
// 				ID: message.To,
// 			}
// 		}
// 	}
// 	for _, message := range groupMessages {
// 		if message.From != userID {
// 			users[message.From] = models.User{
// 				ID: message.From,
// 			}
// 		}
// 	}
// 	for _, message := range groupMessages {
// 		for _, member := range message.Members {
// 			if member != userID {
// 				users[member] = models.User{
// 					ID: member,
// 				}
// 			}
// 		}
// 	}
// 	result := make([]models.User, 0, len(users))
// 	for _, user := range users {
// 		result = append(result, user)
// 	}
// 	return result, nil
// }
