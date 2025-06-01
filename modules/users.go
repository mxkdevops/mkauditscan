package modules

import (
	"os/exec"
	"strings"
)

func GetUsersInfo() map[string]interface{} {
	result := make(map[string]interface{})

	// Get all users from /etc/passwd
	passwdOutput, _ := exec.Command("cut", "-d:", "-f1", "/etc/passwd").Output()
	users := strings.Split(strings.TrimSpace(string(passwdOutput)), "\n")
	result["AllUsers"] = users

	// Get sudo group members
	sudoersOutput, err := exec.Command("getent", "group", "sudo").Output()
	if err == nil {
		fields := strings.Split(strings.TrimSpace(string(sudoersOutput)), ":")
		if len(fields) >= 4 {
			result["SudoUsers"] = strings.Split(fields[3], ",")
		} else {
			result["SudoUsers"] = []string{}
		}
	} else {
		result["SudoUsers"] = "Could not retrieve sudo group"
	}

	// Get last login
	lastlogOutput, err := exec.Command("lastlog").Output()
	if err == nil {
		result["LastLogin"] = string(lastlogOutput)
	} else {
		result["LastLogin"] = "Could not retrieve last login information"
	}

	return result
}
