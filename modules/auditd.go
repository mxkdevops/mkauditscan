package modules

import (
	"os/exec"
	"strings"
)

func GetAuditdStatus() map[string]interface{} {
	result := make(map[string]interface{})

	status, err := exec.Command("auditctl", "-s").Output()
	if err == nil {
		result["AuditdStatus"] = strings.Split(string(status), "\n")
	} else {
		result["AuditdStatus"] = "Could not retrieve auditd status"
	}

	return result
}
