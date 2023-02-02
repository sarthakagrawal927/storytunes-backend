package postgres

import (
	"strconv"

	"github.com/storyTunes/backend.git/cmd/utils"
)

func CreateUserDB(name string, age string, email string) User {
	ageInt, err := strconv.ParseUint(age, 10, 8)
	utils.HandleError(err)
	user := &User{
		Name:  name,
		Age:   uint8(ageInt),
		Email: email,
	}
	PostgresDB.Create(user)
	return *user
}
