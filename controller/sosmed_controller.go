package controller

import (
	"github.com/arifwidiasan/api-taut/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) GetSosmedByUsernameController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	sosmed, err := ce.Svc.GetSosmedByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success get sosmed " + username,
		"data":     sosmed,
	})
}

func (ce *EchoController) UpdateSosmedByUsernameController(c echo.Context) error {
	sosmed := model.Sosmed{}
	if err := c.Bind(&sosmed); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	err := ce.Svc.UpdateSosmedByUsernameService(username, sosmed)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success update sosmed " + username,
	})
}

func (ce *EchoController) GetSosmedByParamUsernameController(c echo.Context) error {
	username := c.Param("username")
	sosmed, err := ce.Svc.GetSosmedByUsernameService(username)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success get sosmed " + username,
		"data":     sosmed,
	})
}
