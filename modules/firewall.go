package modules

import (
	"os/exec"
)

func CheckFirewallStatus() string {
	output := "\n--- Firewall Status ---\n"
	ufwStatus, err := exec.Command("ufw", "status").Output()
	if err == nil {
		output += string(ufwStatus)
	} else {
		iptables, err := exec.Command("iptables", "-L").Output()
		if err == nil {
			output += string(iptables)
		} else {
			output += "Unable to determine firewall status"
		}
	}
	return output
}
