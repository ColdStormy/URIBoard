package web

import (
	"fmt"
	"uriboard/system"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func StartServer() {

	config := system.GetConfiguration()

	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte(config.SecretKey))))
	defineRoutes(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))

}

func defineRoutes(e *echo.Echo) {

	e.GET("/", MainPage)

	e.POST("/login", Login)
}
