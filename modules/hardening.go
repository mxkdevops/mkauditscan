package modules

import (
	"os/exec"
	"strings"
)

func CheckHardening() map[string]interface{} {
	result := make(map[string]interface{})

	checks := map[string]string{
		"Password max days":    "chage -l root | grep 'Maximum' | awk -F: '{print $2}'",
		"Password min days":    "chage -l root | grep 'Minimum' | awk -F: '{print $2}'",
		"World writable files": "find / -xdev -type f -perm -0002 2>/dev/null | wc -l",
		"Hidden processes":     "ps -ef | grep -vE 'ps|grep' | wc -l",
	}

	for label, cmdStr := range checks {
		cmd := exec.Command("bash", "-c", cmdStr)
		out, _ := cmd.Output()
		result[label] = strings.TrimSpace(string(out))
	}

	return result
}
