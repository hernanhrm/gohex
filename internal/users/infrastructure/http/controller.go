package http

import (
	"gohex/internal/request"
	"gohex/internal/users/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type controller struct {
	useCase UseCase
}

func New(useCase UseCase) controller {
	return controller{useCase: useCase}
}

func (c controller) Create(ctx echo.Context) error {
	var m domain.User
	if err := ctx.Bind(&m); err != nil {
		return err
	}

	return c.useCase.Create(ctx.Request().Context(), m)
}

func (c controller) Update(ctx echo.Context) error {
	var m domain.User
	if err := ctx.Bind(&m); err != nil {
		return err
	}

	return c.useCase.Update(ctx.Request().Context(), m)
}

func (c controller) Delete(ctx echo.Context) error {
	id, err := request.GetUUID("id", ctx)
	if err != nil {
		return err
	}

	return c.useCase.Delete(ctx.Request().Context(), id)
}

func (c controller) GetAll(ctx echo.Context) error {
	data, err := c.useCase.GetAll(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

func (c controller) GetByID(ctx echo.Context) error {
	id, err := request.GetUUID("id", ctx)
	if err != nil {
		return err
	}

	data, err := c.useCase.GetByID(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}
