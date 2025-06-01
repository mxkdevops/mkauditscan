package modules

import (
	"os/exec"
	"strings"
)

func CheckRunningServices() map[string]interface{} {
	result := make(map[string]interface{})

	out, _ := exec.Command("systemctl", "list-units", "--type=service", "--state=running").Output()
	services := strings.Split(string(out), "\n")

	risky := []string{}
	for _, svc := range services {
		if strings.Contains(svc, "telnet") || strings.Contains(svc, "ftp") || strings.Contains(svc, "rlogin") {
			risky = append(risky, svc)
		}
	}

	result["Running Services"] = strings.Join(services, "\n")
	if len(risky) > 0 {
		result["Risky Services"] = strings.Join(risky, "\n")
	} else {
		result["Risky Services"] = "✅ No known risky services detected"
	}

	return result
}
