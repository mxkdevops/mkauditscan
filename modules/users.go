// users.go
package modules

import (
	"os/exec"
)

func GetUsersInfo() string {
	output := "\n--- Users Info ---\n"
	passwd, _ := exec.Command("cut", "-d:", "-f1", "/etc/passwd").Output()
	output += "Users:\n" + string(passwd)

	sudoers, _ := exec.Command("getent", "group", "sudo").Output()
	output += "\nSudo Access:\n" + string(sudoers)

	lastlog, _ := exec.Command("lastlog").Output()
	output += "\nLast Login:\n" + string(lastlog)
	return output
}
