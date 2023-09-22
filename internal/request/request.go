package request

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

func GetUUID(name string, c echo.Context) (uuid.UUID, error) {
	fieldStr, ok := c.Get(name).(string)
	if !ok {
		return uuid.UUID{}, fmt.Errorf("request.GetUUID(): field is not of type string")
	}

	field, err := uuid.Parse(fieldStr)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("request.uuid.Parse(): the value %s is not a valid UUID", fieldStr)
	}

	return field, nil
}
