package main

import (
	"database/sql"
	"github.com/labstack/echo/v4"
	"net/http"
	"phonebook/authentication"
	"phonebook/db"
	"phonebook/model"
	"phonebook/store"
)

var sotoonDb *sql.DB


func main() {
	sotoonDb = db.New()
	e := echo.New()
	e.POST("/register", register)
	e.POST("/login", login)

	e.Logger.Fatal(e.Start(":8080"))
}

func register(c echo.Context) error {
	var newUser model.User

	if err := c.Bind(&newUser); err != nil {
		return err
	}

	store.Save(sotoonDb, newUser)

	return c.JSON(http.StatusCreated, newUser)
}

func login(c echo.Context) error {
	var user model.User

	if err := c.Bind(&user); err != nil {
		return err
	}

	if user.Email == "" || user.Password == "" {
		return echo.NewHTTPError(http.StatusBadRequest, "Email and password cannot be empty")
	}

	us, err := store.Retrieve(sotoonDb, user)
	if err != nil {
		return c.JSON(http.StatusNotFound, err.Error())
	}

	token, err := authentication.CreateToken(us.Email)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, token)
}
