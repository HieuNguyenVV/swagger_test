package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"swagger_test/cmd/api/model"
)

type userHandler struct {
	echo *echo.Group
}

func NewUserHandler(echo *echo.Group) *userHandler {
	return &userHandler{echo: echo}
}

// GetUserById godoc
//
// @Id  GetUserById
// @Summary GetUserById
// @Description Get user by id
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param id path string true "id"
// @Success 200 {object} model.Response{}
// @Failure 400 {object} model.AppError{}
// @Router /get/{id} [get]
func (h *userHandler) GetUserById(c echo.Context) error {
	id := c.Param("id")
	fmt.Printf("Inside func GetUserById, ID: %+v", id)
	return c.JSON(http.StatusOK, model.Response{
		Code:   200,
		Status: "Pass",
		Data:   "",
	})
}

// CreateNewUser godoc
//
// @Id 	CreateNewUser
// @Summary			Create New User
// @Description		Save user data in Db.
// @Tags			abc
// @Accept          application/json
// @Produce			application/json
// @Param			user body model.CreateUserRequest  true "Create user"
// @Success			200 {object} model.Response{}
// @Failure			400	{object}	model.AppError{}
// @Failure			404	{object}	model.AppError{}
// @Failure			500	{object}	model.AppError{}
// @Router			/create [post]
func (h *userHandler) CreateNewUser(c echo.Context) error {
	var req model.CreateUserRequest
	if err := c.Bind(&req); err != nil {
		fmt.Printf("BAD REQUEST")
		return c.JSON(http.StatusBadRequest, model.AppError{
			StatusCode: 0,
			RootErr:    nil,
			Message:    "",
			Log:        "",
			Key:        "",
		})
	}
	fmt.Printf("Inside func create new user, value of request: %+v", req)
	return c.JSON(http.StatusOK, model.Response{
		Code:   200,
		Status: "Pass",
		Data:   req,
	})
}
