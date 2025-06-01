package modules

import (
	"os/exec"
	"strings"
)

func AuditSSHConfig() map[string]interface{} {
	result := make(map[string]interface{})

	// Read sshd_config
	conf, _ := exec.Command("cat", "/etc/ssh/sshd_config").Output()
	lines := strings.Split(string(conf), "\n")

	rules := map[string]string{
		"PermitRootLogin":        "no",
		"PasswordAuthentication": "no",
		"X11Forwarding":          "no",
		"PermitEmptyPasswords":   "no",
		"UsePAM":                 "yes",
	}

	for key, expected := range rules {
		found := false
		for _, line := range lines {
			if strings.HasPrefix(line, key) {
				found = true
				if strings.Contains(line, expected) {
					result[key] = "✅ " + line
				} else {
					result[key] = "❌ " + line + " (Expected: " + expected + ")"
				}
				break
			}
		}
		if !found {
			result[key] = "⚠️ Not set (Expected: " + expected + ")"
		}
	}

	return result
}
