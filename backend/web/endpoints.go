package web

import (
	"fmt"
	"net/http"
	"uriboard/system"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func MainPage(c echo.Context, config system.Configuration_t) error {
	sess, _ := session.Get("session", c)

	name := "World"
	if true == sess.Values["loggedIn"] {
		name = "Admin"
	}

	msg := fmt.Sprintf("Hello, %s!", name)
	return c.String(http.StatusOK, msg)
}

func Login(c echo.Context, config system.Configuration_t) error {
	password := c.QueryParam("password")

	if password == config.AdminPassword {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}
		sess.Values["loggedIn"] = true
		sess.Save(c.Request(), c.Response())
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusUnauthorized)
}
