package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

type Configuration struct {
	AdminPassword string `json:"admin_password"`
	Port          int    `json:"port"`
	SecretKey     string `json:"secret_key"`
}

var config Configuration

func loadConfiguration() {
	jsonFile, err := os.Open("config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	byteContent, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal(byteContent, &config)
	fmt.Println(config)

}

func main() {

	loadConfiguration()

	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("secret"))))
	e.GET("/", func(c echo.Context) error {
		sess, _ := session.Get("session", c)

		name := "World"
		if true == sess.Values["loggedIn"] {
			name = "Admin"
		}

		msg := fmt.Sprintf("Hello, %s!", name)
		return c.String(http.StatusOK, msg)
	})

	e.POST("/login", func(c echo.Context) error {
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
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.Port)))
}
