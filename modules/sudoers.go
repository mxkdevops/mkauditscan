package modules

import (
	"os/exec"
	"strings"
)

func GetSudoersInfo() map[string]interface{} {
	result := make(map[string]interface{})

	// Get list of users with sudo privileges (from /etc/sudoers)
	sudoersFile, err := exec.Command("grep", "-v", "^#", "/etc/sudoers").Output()
	if err == nil {
		lines := strings.Split(strings.TrimSpace(string(sudoersFile)), "\n")
		result["SudoersFile"] = lines
	} else {
		result["SudoersFile"] = "Could not read /etc/sudoers"
	}

	// Check sudo group members using getent
	sudoGroup, err := exec.Command("getent", "group", "sudo").Output()
	if err == nil {
		fields := strings.Split(strings.TrimSpace(string(sudoGroup)), ":")
		if len(fields) >= 4 {
			users := strings.Split(fields[3], ",")
			result["SudoGroupMembers"] = users
		} else {
			result["SudoGroupMembers"] = []string{}
		}
	} else {
		result["SudoGroupMembers"] = "Could not retrieve sudo group members"
	}

	// Optionally check if any users have passwordless sudo
	pwless, err := exec.Command("grep", "NOPASSWD", "/etc/sudoers").Output()
	if err == nil {
		lines := strings.Split(strings.TrimSpace(string(pwless)), "\n")
		result["PasswordlessSudo"] = lines
	} else {
		result["PasswordlessSudo"] = "None found or inaccessible"
	}

	return result
}
