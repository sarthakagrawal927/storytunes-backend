package server

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func StartServer() {
	e := echo.New()
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.GET("/", testServer)

	// user routes
	e.POST("/users/create", createUser)

	e.Logger.Fatal(e.Start(":1323"))
}
