package factories

import (
	userRepo "github.com/nugrohosam/gocashier/repositories/user"
)

// CreateUser ...
func CreateUser() userRepo.User {
	user := userRepo.User{
		Name:     "asu",
		Username: "asu",
		Password: "asu",
		Email:    "asu@asu.com",
	}

	return user
}
