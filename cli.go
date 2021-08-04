package sbutils

import (
	"bufio"
	"bytes"
	"errors"
	"fmt"
	"os/exec"
	"strings"
)

// ExecuteCommand : Execute Shell Command
func ExecuteCommand(input string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", input)
	stdout, stderr := cmd.Output()
	return strings.TrimSpace(string(stdout)), stderr
}

// ExecuteCommand : Execute Shell Command With Error Buffer(more detailed errors)
func ExecuteCommandWithErrorBuffer(command string) (string, error) {
	cmd := exec.Command("/bin/bash", "-c", command)
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		return "", errors.New(stderr.String())
	}

	return strings.TrimSuffix(string(out.String()), "\n"), nil
}

// ExecuteWithLiveOutput : executing given command with realtime output
func ExecuteWithLiveOutput(command string) error {
	cmd := exec.Command("/bin/bash", "-c", command)
	stdout, _ := cmd.StdoutPipe()
	var stderr bytes.Buffer
	cmd.Stderr = &stderr

	_ = cmd.Start()
	scanner := bufio.NewScanner(stdout)
	for scanner.Scan() {
		m := scanner.Text()
		fmt.Println(m)
	}
	_ = cmd.Wait()

	if stderr.Len() > 0 {
		return errors.New(stderr.String())
	}

	return nil
}
