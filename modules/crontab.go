package modules

import (
	"os/exec"
	"strings"
)

func GetCrontabEntries() map[string]interface{} {
	result := make(map[string]interface{})

	// Get list of users from /etc/passwd
	passwdOutput, err := exec.Command("cut", "-d:", "-f1", "/etc/passwd").Output()
	if err != nil {
		result["Error"] = "Failed to retrieve user list"
		return result
	}
	users := strings.Split(strings.TrimSpace(string(passwdOutput)), "\n")

	// Get root/system crontab from /etc/crontab
	systemCrontab, err := exec.Command("cat", "/etc/crontab").Output()
	if err == nil {
		result["SystemCrontab"] = string(systemCrontab)
	} else {
		result["SystemCrontab"] = "Could not read /etc/crontab"
	}

	// Get each user's crontab
	userCrontabs := make(map[string]string)
	for _, user := range users {
		output, err := exec.Command("crontab", "-l", "-u", user).Output()
		if err == nil {
			userCrontabs[user] = string(output)
		}
	}
	result["UserCrontabs"] = userCrontabs

	// Optional: Include cron.d entries
	cronDOutput, err := exec.Command("ls", "-1", "/etc/cron.d").Output()
	if err == nil {
		files := strings.Split(strings.TrimSpace(string(cronDOutput)), "\n")
		cronDEntries := make(map[string]string)
		for _, file := range files {
			entry, err := exec.Command("cat", "/etc/cron.d/"+file).Output()
			if err == nil {
				cronDEntries[file] = string(entry)
			}
		}
		result["CronD"] = cronDEntries
	}

	return result
}
