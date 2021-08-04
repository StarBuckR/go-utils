package utils

import (
	"os/user"
)

func IsRoot() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, err
	}

	return currentUser.Username == "root", nil
}
