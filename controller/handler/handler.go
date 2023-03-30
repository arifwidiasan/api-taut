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
	api.POST("/admins", cont.CreateAdminController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.GET("/admins", cont.GetAllAdminController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/admins/:id", cont.GetAdminByIDController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.PUT("/admins/:id", cont.UpdateAdminController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.DELETE("/admins/:id", cont.DeleteAdminController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.POST("/admins/login", cont.LoginAdminController)
	api.POST("/admins/changepass", cont.ChangePassAdminController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.POST("/admins/users", cont.AdminCreateUserController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.GET("/admins/users", cont.AdminGetAllUserController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/admins/users/:id", cont.AdminGetUserByIDController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.PUT("/admins/users/:id", cont.AdminUpdateUserByIDController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.DELETE("/admins/users/:id", cont.AdminDeleteUserByIDController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.POST("/users", cont.CreateUserController)
	api.GET("/users", cont.GetUserByUsernameController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.PUT("/users", cont.UpdateUserByUsernameController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.POST("/users/login", cont.LoginUserController)
	api.POST("/users/changepass", cont.ChangePassUserController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/sosmeds", cont.GetSosmedByUsernameController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.PUT("/sosmeds", cont.UpdateSosmedByUsernameController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.POST("/sections", cont.CreateSectionController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.GET("/sections", cont.GetAllSectionByUserIDController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/sections/:id", cont.GetOneSectionByUserIDandIDController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.PUT("/sections/:id", cont.UpdateSectionByUserIDandIDController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.DELETE("/sections/:id", cont.DeleteSectionByUserIDandIDController, middleware.JWT([]byte(conf.JWT_KEY)))

	api.GET("/profiles/users/:username", cont.GetUserByParamUsernameController)
	api.GET("/profiles/sosmeds/:username", cont.GetSosmedByParamUsernameController)
	api.GET("/profiles/sections/:username", cont.GetAllSectionByParamUsernameController)
	api.GET("/profiles/users/:username/photo", cont.GetProfilePictureController)
	api.GET("/profiles/users/:username/qrcode", cont.GetQRCodeController)

	api.POST("/uploads/profile-picture", cont.UploadProfilePictureController, middleware.JWT([]byte(conf.JWT_KEY)))
	api.DELETE("/uploads/profile-picture", cont.DeleteProfilePictureController, middleware.JWT([]byte(conf.JWT_KEY)))

}
