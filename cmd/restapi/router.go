package main

import (
	users "gohex/internal/users/infrastructure/http"
)

func initRoutes() {
	users.NewRouter()
}
