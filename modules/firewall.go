package modules

import (
	"os/exec"
)

func CheckFirewallStatus() map[string]interface{} {
	result := make(map[string]interface{})
	result["FirewallCheck"] = "Started"

	// Try ufw
	ufwStatus, err := exec.Command("ufw", "status").Output()
	if err == nil {
		result["Tool"] = "ufw"
		result["Status"] = string(ufwStatus)
		return result
	}

	// Fallback to iptables
	iptablesStatus, err := exec.Command("iptables", "-L").Output()
	if err == nil {
		result["Tool"] = "iptables"
		result["Status"] = string(iptablesStatus)
		return result
	}

	// If both failed
	result["Tool"] = "none"
	result["Status"] = "Unable to determine firewall status"
	return result
}
