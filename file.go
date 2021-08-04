package sbutils

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
	defer file.Close()

	file.WriteString(content)
	return nil
}

// Delete all file contents and rewrite it
func WriteToFileTruncate(filePath string, content string) error {
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	file.WriteString(content)
	return nil
}
