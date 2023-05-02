package controller

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/arifwidiasan/api-taut/model"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"github.com/xuri/excelize/v2"
)

func (ce *EchoController) AdminCreateUserController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not an admin",
		})
	}

	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.AdminCreateUserService(user)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success create user " + user.Username,
	})
}

func (ce *EchoController) AdminGetAllUserController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not an admin",
		})
	}

	users := ce.Svc.AdminGetAllUserService()

	return c.JSON(200, map[string]interface{}{
		"messages": "success get all user",
		"data":     users,
	})
}

func (ce *EchoController) AdminGetUserByIDController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not an admin",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	user, err := ce.Svc.AdminGetUserByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	sosmed, _ := ce.Svc.GetSosmedByUserIDService(id_int)

	return c.JSON(200, map[string]interface{}{
		"messages": "success",
		"data":     user,
		"sosmed":   sosmed,
	})
}

func (ce *EchoController) AdminUpdateUserByIDController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not an admin",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	user := model.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	err = ce.Svc.AdminUpdateUserByIDService(id_int, user)
	if err != nil {
		return c.JSON(500, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success update user",
	})
}

func (ce *EchoController) AdminDeleteUserByIDController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not an admin",
		})
	}

	id := c.Param("id")
	id_int, _ := strconv.Atoi(id)
	err = ce.Svc.AdminDeleteUserByIDService(id_int)
	if err != nil {
		return c.JSON(404, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	return c.JSON(200, map[string]interface{}{
		"messages": "success delete user",
	})
}

func (ce *EchoController) AdminDownloadTemplateController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not an admin",
		})
	}

	return c.Attachment("helper/template-taut.xlsx", "template-taut.xlsx")
}

func (ce *EchoController) AdminInsertBatchUserController(c echo.Context) error {
	username := ce.Svc.ClaimToken(c.Get("user").(*jwt.Token))
	_, err := ce.Svc.GetAdminByUsernameService(username)
	if err != nil {
		return c.JSON(403, map[string]interface{}{
			"messages": "forbidden, not an admin",
		})
	}

	//upload file excel
	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": err.Error(),
		})
	}

	filebyte, _ := io.ReadAll(src)
	filetype := http.DetectContentType(filebyte)
	if filetype != "application/zip" {
		return c.JSON(400, map[string]interface{}{
			"messages": "file type not allowed",
		})
	}

	defer src.Close()

	filename := "temp.xlsx"
	_ = os.WriteFile("../uploads/excel/"+filename, filebyte, 0777)
	/* if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error write file",
		})
	} */

	//read file excel
	filepath := "../uploads/excel/" + filename
	excel, err := excelize.OpenFile(filepath)
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error open file",
		})
	}

	//check template
	cols, err := excel.GetCols("Sheet1")
	if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error read column file",
		})
	}

	var check string
	for _, col := range cols {
		check = fmt.Sprintf("%s%s", check, col[0])
	}

	if check != "nonamajabatan/pekerjaanno.telpemailusername" {
		return c.JSON(400, map[string]interface{}{
			"messages": "template salah",
		})
	}

	//insert batch
	dimension, _ := excel.GetSheetDimension("Sheet1")
	rangeRow := strings.SplitAfter(dimension, "F")
	rangeRowInt, _ := strconv.Atoi(rangeRow[1])

	//set cell status and password
	excel.SetCellValue("Sheet1", "G1", "password")
	excel.SetCellValue("Sheet1", "H1", "status")
	excel.SetCellValue("Sheet1", "I1", "keterangan")

	for i := 2; i <= rangeRowInt; i++ {
		//set user
		user := model.User{}
		user.Name, _ = excel.GetCellValue("Sheet1", "B"+strconv.Itoa(i))
		user.Job, _ = excel.GetCellValue("Sheet1", "C"+strconv.Itoa(i))
		user.PhoneNumber, _ = excel.GetCellValue("Sheet1", "D"+strconv.Itoa(i))
		user.Email, _ = excel.GetCellValue("Sheet1", "E"+strconv.Itoa(i))
		user.Username, _ = excel.GetCellValue("Sheet1", "F"+strconv.Itoa(i))
		user.BornDate = time.Now()

		//insert and set user
		err = ce.Svc.AdminCreateUserService(user)
		if err != nil {
			excel.SetCellValue("Sheet1", "H"+strconv.Itoa(i), "gagal")
			excel.SetCellValue("Sheet1", "I"+strconv.Itoa(i), err.Error())
		} else {
			excel.SetCellValue("Sheet1", "G"+strconv.Itoa(i), user.Username+"2023")
			excel.SetCellValue("Sheet1", "H"+strconv.Itoa(i), "sukses")
		}
	}

	//save file
	_ = excel.SaveAs("../uploads/excel/hasil_insert.xlsx")
	/* if err != nil {
		return c.JSON(400, map[string]interface{}{
			"messages": "error save file result",
		})
	} */

	defer excel.Close()

	return c.Attachment("../uploads/excel/hasil_insert.xlsx", "hasil_insert.xlsx")
}
