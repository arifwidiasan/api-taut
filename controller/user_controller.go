package controller

import (
	"github.com/arifwidiasan/api-taut/model"
	"github.com/golang-jwt/jwt"
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

func (ce *EchoController) LoginUserController(c echo.Context) error {
	userLogin := model.UserLogin{}
	if err := c.Bind(&userLogin); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	token, statusCode := ce.Svc.LoginUserService(userLogin.Username, userLogin.Password)
	switch statusCode {
	case 401:
		return c.JSON(401, map[string]interface{}{
			"messages": "username atau password salah",
		})

	case 500:
		return c.JSON(500, map[string]interface{}{
			"messages": "internal, error create token",
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success login as " + userLogin.Username,
		"token":    token,
	})
}

func (ce *EchoController) GetUserByUsernameController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	user, err := ce.Svc.GetUserByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success get user " + username,
		"data":     user,
	})
}
