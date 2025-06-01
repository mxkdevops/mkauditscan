package modules

import (
	"os/exec"
	"strings"
)

func GetAuthInfo() map[string]interface{} {
	result := make(map[string]interface{})

	authLog, err := exec.Command("tail", "-n", "100", "/var/log/auth.log").Output()
	if err == nil {
		result["RecentAuthLog"] = strings.Split(string(authLog), "\n")
	} else {
		result["RecentAuthLog"] = "Could not retrieve /var/log/auth.log"
	}

	return result
}
