package controller

import (
	"strconv"

	"github.com/arifwidiasan/api-taut/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (ce *EchoController) CreateSectionController(c echo.Context) error {
	section := model.Section{}
	if err := c.Bind(&section); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	err := ce.Svc.CreateSectionService(username, section)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success create section",
	})
}

func (ce *EchoController) GetAllSectionByUserIDController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	sections := ce.Svc.GetAllSectionByUserIDService(username)

	return c.JSON(200, map[string]interface{}{
		"messages": "success get all section from " + username,
		"data":     sections,
	})
}

func (ce *EchoController) GetOneSectionByUserIDandIDController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)

	section, err := ce.Svc.GetOneSectionByUserIDandIDService(username, id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success get section from " + username,
		"data":     section,
	})
}

func (ce *EchoController) UpdateSectionByUserIDandIDController(c echo.Context) error {
	section := model.Section{}
	if err := c.Bind(&section); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)

	err := ce.Svc.UpdateSectionByUserIDandIDService(username, id_int, section)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success update section from " + username,
	})
}
