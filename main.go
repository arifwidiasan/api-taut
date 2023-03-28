package main

import (
	conf "github.com/arifwidiasan/api-taut/config"
	handler "github.com/arifwidiasan/api-taut/controller/handler"
	"github.com/arifwidiasan/api-taut/helper"

	"github.com/labstack/echo/v4"
)

func main() {
	helper.CreateFolder("../uploads")
	helper.CreateFolder("../uploads/profile-picture")
	config := conf.InitConfiguration()
	e := echo.New()

	handler.RegisterGroupAPI(e, config)

	e.Logger.Fatal(e.Start((config.SERVER_ADDRESS)))
}
