package web

import (
	"fmt"
	"uriboard/system"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func StartServer(config system.Configuration_t) {

	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(config.SecretKey))))
	defineRoutes(e, config)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}

func defineRoutes(e *echo.Echo, config system.Configuration_t) {

	e.GET("/", func(c echo.Context) error { return MainPage(c, config) })

	e.POST("/login", func(c echo.Context) error { return Login(c, config) })
}
