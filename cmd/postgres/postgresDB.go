package postgres

import (
	"github.com/storyTunes/backend.git/cmd/utils"
)

func CreateUserDB(name string, email string) User {
	utils.HandleError(err)
	user := &User{
		Name: name,

		Email: email,
	}
	PostgresDB.Create(user)
	return *user
}
