package controller

import (
	"github.com/arifwidiasan/api-taut/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateUserController(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err := ce.Svc.CreateUserService(user)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success create user " + user.Username,
	})
}
