package main

import (
	"github.com/labstack/echo/v4"
	"swagger_test/cmd/api/handler"
	_ "swagger_test/cmd/docs"
)

// @title 	Tag Service API
// @version	1.0
// @description A Tag service API in Go using Echo framework
// @termsOfService http://swagger.io/terms/

// @contact.name  API Support
// @contact.url   http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url  http://www.apache.org/licenses/LICENSE-2.0.html

// @host     localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apiKey Bearer
// @in                         header
// @name                       Authorization
// @description                Enter the token with the `Bearer: ` prefix, e.g. "Bearer abcde12345"
func main() {
	e := echo.New()
	baseRouter := e.Group("/api/v1")

	h := handler.NewUserHandler(baseRouter)
	h.MapRouter()
	e.Logger.Fatal(e.Start(":8080"))
}
