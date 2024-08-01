package scripts

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

func GetPythonGestureInput(prompt string) (int, error) {
	fmt.Println(prompt)
	cmd := exec.Command("python3", "scripts/gesture_recognition.py")
	output, err := cmd.Output()
	if err != nil {
		fmt.Println("Error executing Python script:", err)
		return -1, err
	}

	result, err := strconv.Atoi(strings.TrimSpace(string(output)))
	if err != nil {
		fmt.Println("Error converting Python script output:", err)
		return -1, err
	}

	return result, err
}
