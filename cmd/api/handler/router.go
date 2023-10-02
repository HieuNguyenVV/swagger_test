package handler

import echoSwagger "github.com/swaggo/echo-swagger"

func (h *userHandler) MapRouter() {
	h.echo.GET("/docs/*any", echoSwagger.WrapHandler)
	h.echo.GET("/get/:id", h.GetUserById)
	h.echo.POST("/create", h.CreateNewUser)
}
