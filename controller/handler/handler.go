package handler

import (
	"github.com/arifwidiasan/api-taut/config"
	"github.com/arifwidiasan/api-taut/controller"
	"github.com/arifwidiasan/api-taut/database"

	m "github.com/arifwidiasan/api-taut/middleware"
	"github.com/arifwidiasan/api-taut/repository"
	"github.com/arifwidiasan/api-taut/service"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func RegisterGroupAPI(e *echo.Echo, conf config.Config) {

	db := database.InitDB(conf)
	repo := repository.NewMysqlRepository(db)

	svc := service.NewService(repo, conf)

	cont := controller.EchoController{
		Svc: svc,
	}

	e.GET("/v1/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	e.POST("/v1/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{
			"message": "your request awesome",
		})
	})

	api := e.Group("/v1", middleware.CORS())

	m.LogMiddleware(e)
	api.POST("/admins/login", cont.LoginAdminController)

}
