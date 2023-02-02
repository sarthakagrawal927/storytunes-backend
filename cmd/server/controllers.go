package server

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/storyTunes/backend.git/cmd/postgres"
	"github.com/storyTunes/backend.git/cmd/utils"
)

func testServer(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}

func createUser(c echo.Context) error {
	defer utils.Recovery()
	user := postgres.CreateUserDB(c.FormValue("name"), c.FormValue("age"), c.FormValue("email"))
	return c.JSON(http.StatusOK, user)
}
