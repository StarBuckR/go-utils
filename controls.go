package utils

import (
	"fmt"
	"os/user"
	"strings"
)

// Check if we ran as root
func IsRoot() (bool, error) {
	currentUser, err := user.Current()
	if err != nil {
		return false, err
	}

	return currentUser.Username == "root", nil
}

func CheckIsAppAlreadyRunning(appName string) (bool, error) {
	out, err := ExecuteCommand(fmt.Sprintf("ps aux | grep %s | grep -v grep | grep -v sudo", appName))
	if err != nil {
		return false, err
	}

	if len(strings.Split(out, "\n")) > 1 {
		return true, nil
	}

	return false, nil
}
