package config

import (
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/labstack/echo/v4"
)

type Router struct {
	Logger       Logger
	DBPool       *pgxpool.Pool
	EchoHttp     *echo.Echo
	RemoteConfig RemoteConfig
}
