package internal

import "gohex/internal/users"

// Load starts the build and load process of every domain dependency
func Load() {
	users.BuildAndLoad()
}
