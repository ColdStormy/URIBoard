package web

import (
	"fmt"
	"net/http"
	"uriboard/usermanager"

	"github.com/google/uuid"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func MainPage(c echo.Context) error {
	sess, _ := session.Get("session", c)
	name := "World"

	uid_interface, found := sess.Values["userId"]

	if found {
		uid_string := fmt.Sprintf("%v", uid_interface)

		if userId, err := uuid.Parse(uid_string); err == nil {
			if User, found := usermanager.GetUserById(userId); found {
				name = User.Name
			}

		}
	}

	msg := fmt.Sprintf("Hello, %s!", name)
	return c.String(http.StatusOK, msg)
}

func Login(c echo.Context) error {
	type QueryParams struct {
		Username string `json:"username" form:"username" query:"username" param:"username"`
		Password string `json:"password" form:"password" query:"password" param:"password"`
	}

	params := new(QueryParams)
	if err := c.Bind(params); err != nil {
		return err
	}

	if usermanager.ValidateUser(params.Username, params.Password) {
		sess, _ := session.Get("session", c)
		sess.Options = &sessions.Options{
			Path:     "/",
			MaxAge:   86400 * 7,
			HttpOnly: true,
		}

		sess.Values["userId"] = usermanager.NameToId(params.Username).String()
		sess.Save(c.Request(), c.Response())
		return c.NoContent(http.StatusOK)
	}

	return c.NoContent(http.StatusUnauthorized)
}
