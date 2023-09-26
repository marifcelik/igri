package db

import (
	"context"

	"go-chat/pkg/models"
	"go-chat/pkg/utils"

	"github.com/charmbracelet/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var db *gorm.DB

type loginPartialUser struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func init() {
	var err error
	db, err = gorm.Open(postgres.Open(utils.GetDBURL()), &gorm.Config{
		TranslateError: true,
		Logger:         logger.Default,
	})
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(
		&models.User{},
		&models.UserMessage{},
		&models.Group{},
		&models.GroupMessage{},
		&models.GroupMessageSeenBy{},
	)
}

func GetUserByID(id string, c context.Context) (models.User, error) {
	user := models.User{}
	result := db.WithContext(c).First(&user)
	return user, result.Error
}

func GetUserByUsername(u string, c context.Context) (loginPartialUser, error) {
	partialUser := loginPartialUser{}
	result := db.WithContext(c).Where("username = ?", u).First(&models.User{}).Scan(&partialUser)
	return partialUser, result.Error
}

func CreateUser(u *models.User, c context.Context) (int, error) {
	result := db.WithContext(c).Create(u)
	return int(result.RowsAffected), result.Error
}

func CheckUsername(u string, c context.Context) (bool, error) {
	result := db.WithContext(c).
		Select("id").
		Where("username = ?", u).
		First(&models.User{})
	return result.RowsAffected > 0, result.Error
}
