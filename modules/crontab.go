package modules

import (
	"os/exec"
)

func GetCrontabEntries() string {
	output := "\n--- Crontab Entries ---\n"
	etcCrontab, _ := exec.Command("cat", "/etc/crontab").Output()
	output += "\n/etc/crontab:\n" + string(etcCrontab)

	userCrontab, _ := exec.Command("crontab", "-l").Output()
	output += "\nUser Crontab:\n" + string(userCrontab)
	return output
}
