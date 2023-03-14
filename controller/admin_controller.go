package controller

import (
	"net/http"
	"strconv"

	"github.com/arifwidiasan/api-taut/model"
	"github.com/golang-jwt/jwt"
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
			"messages": "internal, error create token",
		}, "  ")
	}

	return c.JSONPretty(http.StatusOK, map[string]interface{}{
		"messages": "success login as " + adminLogin.Username,
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

	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	err := ce.Svc.ChangePassAdminService(username, adminPass.OldPass, adminPass.NewPass)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success change password admin " + username,
	})
}

func (ce *EchoController) CreateAdminController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not an admin",
		})
	}

	if username != "admin" {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not master admin",
		})
	}

	admin := model.Admin{}

	if err := c.Bind(&admin); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.CreateAdminService(admin)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success create admin " + username,
	})
}

func (ce *EchoController) GetAllAdminController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not an admin",
		})
	}

	admins := ce.Svc.GetAllAdminService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success get all admin",
		"data":     admins,
	})
}

func (ce *EchoController) GetAdminByIDController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not an admin",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	admin, err := ce.Svc.GetAdminByIDService(id_int)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success get admin",
		"data":     admin,
	})
}
