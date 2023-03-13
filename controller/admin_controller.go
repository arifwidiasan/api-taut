package controller

import (
	"net/http"

	"github.com/arifwidiasan/api-taut/model"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) LoginAdminController(c echo.Context) error {
	adminLogin := model.AdminLogin{}

	if err := c.Bind(&adminLogin); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	token, statusCode := ce.Svc.LoginAdmin(adminLogin.Username, adminLogin.Password)
	switch statusCode {
	case http.StatusUnauthorized:
		return c.JSONPretty(http.StatusUnauthorized, map[string]interface{}{
			"messages": "username atau password salah",
		}, "  ")

	case http.StatusInternalServerError:
		return c.JSONPretty(http.StatusInternalServerError, map[string]interface{}{
			"messages": "internal",
		}, "  ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success",
		"token":    token,
	}, "  ")
}

func (ce *EchoController) ChangePassAdminController(c echo.Context) error {
	adminPass := model.AdminChangePass{}

	if err := c.Bind(&adminPass); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err := ce.Svc.ChangePassAdminService(adminPass.OldPass, adminPass.NewPass)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success change password admin",
	})
}
