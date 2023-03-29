package controller

import (
	"io"
	"net/http"
	"os"
	"strings"

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

func (ce *EchoController) ChangePassUserController(c echo.Context) error {
	userPass := model.UserChangePass{}
	if err := c.Bind(&userPass); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	err := ce.Svc.ChangePassUserService(username, userPass.OldPass, userPass.NewPass)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success change password user " + username,
	})
}

func (ce *EchoController) UpdateUserByUsernameController(c echo.Context) error {
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	err := ce.Svc.UpdateUserByUsernameService(username, user)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success update user " + username,
	})
}

func (ce *EchoController) GetUserByParamUsernameController(c echo.Context) error {
	username := c.Param("username")
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

func (ce *EchoController) UploadProfilePictureController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error open file",
		})
	}

	if file.Size >= 2200000 {
		return c.JSON(400, map[string]interface{}{
			"messages": "max file is 2 MB",
		})
	}

	filebyte, _ := io.ReadAll(src)
	filetype := http.DetectContentType(filebyte)
	if !strings.Contains(filetype, "image") {
		return c.JSON(400, map[string]interface{}{
			"messages": "file is not image",
		})
	}

	defer src.Close()

	user := model.User{}
	filename := username + ".png" //+ strings.SplitAfter(filetype, "/")[1]
	user.ProfilePicturePathFile = filename

	err = os.WriteFile("../uploads/profile-picture/"+filename, filebyte, 0777)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error write file",
		})
	}

	_ = ce.Svc.UpdateUserByUsernameService(username, user)

	return c.JSON(200, map[string]interface{}{
		"messages": "success upload profile picture user " + username,
	})
}

func (ce *EchoController) DeleteProfilePictureController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))

	err := os.Remove("../uploads/profile-picture/" + username + ".png")
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error delete file",
		})
	}

	user := model.User{}
	user.ProfilePicturePathFile = " "

	_ = ce.Svc.UpdateUserByUsernameService(username, user)

	return c.JSON(200, map[string]interface{}{
		"messages": "success delete profile picture user " + username,
	})
}

func (ce *EchoController) GetProfilePictureController(c echo.Context) error {
	username := c.Param("username")
	file, err := os.Open("../uploads/profile-picture/" + username + ".png")
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error open file",
		})
	}

	defer file.Close()

	filebyte, _ := io.ReadAll(file)

	return c.Blob(200, "image/png", filebyte)
}
