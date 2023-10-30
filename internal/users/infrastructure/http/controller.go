package http

import (
	"gohex/internal/request"
	"gohex/internal/users/domain"
	"net/http"

	"github.com/labstack/echo/v4"
)

type Controller struct {
	UseCase UseCase
}

func (c Controller) Create(ctx echo.Context) error {
	var m domain.User
	if err := ctx.Bind(&m); err != nil {
		return err
	}

	return c.UseCase.Create(ctx.Request().Context(), m)
}

func (c Controller) Update(ctx echo.Context) error {
	var m domain.User
	if err := ctx.Bind(&m); err != nil {
		return err
	}

	return c.UseCase.Update(ctx.Request().Context(), m)
}

func (c Controller) Delete(ctx echo.Context) error {
	id, err := request.GetUUID("id", ctx)
	if err != nil {
		return err
	}

	return c.UseCase.Delete(ctx.Request().Context(), id)
}

func (c Controller) List(ctx echo.Context) error {
	data, err := c.UseCase.List(ctx.Request().Context())
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}

func (c Controller) Get(ctx echo.Context) error {
	id, err := request.GetUUID("id", ctx)
	if err != nil {
		return err
	}

	data, err := c.UseCase.Get(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, data)
}
