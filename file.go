package utils

import (
	"os"
)

// Check required file or directory by giving its path
func CheckRequiredFileOrDir(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	}

	return true
}

// Read file and get it's string output
func ReadFile(path string) (string, error) {
	b, err := os.ReadFile(path)

	return string(b), err
}

// Read file and get it's string output
func AppendToFile(path string, content string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	file.WriteString(content)
	return nil
}
