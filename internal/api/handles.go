package api

import (
	"net/http"

	"github.com/brianepv1/sv-microservice-login-golang/internal/api/dtos"
	"github.com/brianepv1/sv-microservice-login-golang/internal/service"
	"github.com/labstack/echo/v4"
)

type badRequestResponseMessage struct {
	Message string `json:"message"`
}

func (a *API) RegisterUser(c echo.Context) error {
	ctx := c.Request().Context()
	params := dtos.RegisterUser{}

	err := c.Bind(&params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, badRequestResponseMessage{Message: "Invalid request"})
	}

	err = a.dataValidator.Struct(params)
	if err != nil {
		return c.JSON(http.StatusBadRequest, badRequestResponseMessage{Message: err.Error()})
	}

	err = a.serv.RegisterUser(ctx, params.Email, params.Name, params.Password)
	if err != nil {
		if err == service.ErrUserAlreadyExists {
			return c.JSON(http.StatusConflict, badRequestResponseMessage{Message: "User already exists"})
		}

		return c.JSON(http.StatusInternalServerError, badRequestResponseMessage{Message: "Internal server error: " + err.Error()})
	}

	return c.JSON(http.StatusCreated, nil)
}
