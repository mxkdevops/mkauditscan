package modules

import (
	"os/exec"
	"strings"
)

func GetNetworkInfo() map[string]interface{} {
	result := make(map[string]interface{})

	ifconfig, err := exec.Command("ip", "a").Output()
	if err == nil {
		result["IPAddresses"] = strings.Split(string(ifconfig), "\n")
	} else {
		result["IPAddresses"] = "Could not retrieve IP configuration"
	}

	routes, err := exec.Command("ip", "route").Output()
	if err == nil {
		result["Routes"] = strings.Split(string(routes), "\n")
	} else {
		result["Routes"] = "Could not retrieve routing table"
	}

	return result
}
