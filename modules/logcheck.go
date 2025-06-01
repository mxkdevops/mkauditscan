package modules

import (
	"os/exec"
	"strings"
)

func GetLogCheckResults() map[string]interface{} {
	result := make(map[string]interface{})

	logcheck, err := exec.Command("logcheck").Output()
	if err == nil {
		result["LogCheckOutput"] = strings.Split(string(logcheck), "\n")
	} else {
		result["LogCheckOutput"] = "Logcheck not found or not installed"
	}

	return result
}
