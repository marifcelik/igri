package db

import (
	"go-chat/models"
	projectUtils "go-chat/utils"

	"github.com/surrealdb/surrealdb.go"
)

var DB *surrealdb.DB

func init() {
	var err error
	DB, err = surrealdb.New(projectUtils.GetDBConnStr())
	if err != nil {
		projectUtils.CheckErr(projectUtils.ErrMsg{Err: err, Msg: "db connection error"})
	}

	_, err = DB.Signin(map[string]interface{}{
		"user": "root",
		"pass": "root",
	})
	if err != nil {
		projectUtils.CheckErr(err)
	}

	_, err = DB.Use("test", "test")
	if err != nil {
		projectUtils.CheckErr(err)
	}

	user := models.User{
		Name:    "arif",
		Surname: "arifcelik",
	}

	data, err := DB.Create("user", user)
	if err != nil {
		projectUtils.CheckErr(err)
	}

}
