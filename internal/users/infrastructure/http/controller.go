package http

import (
	"gohex/internal/users/dto"

	"github.com/labstack/echo/v4"
)

type controller struct {
	useCase UseCase
}

func New(useCase UseCase) controller {
	return controller{useCase: useCase}
}

func (c controller) Create(ctx echo.Context) error {
	var createDto dto.Create
	if err := ctx.Bind(&createDto); err != nil {
		return err
	}

	return c.useCase.Create(ctx.Request().Context(), createDto)
}
