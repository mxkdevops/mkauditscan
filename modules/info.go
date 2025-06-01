package modules

import (
	"os/exec"
	"strings"
)

func GetSystemInfo() map[string]interface{} {
	result := make(map[string]interface{})

	hostname, _ := exec.Command("hostname").Output()
	osRelease, _ := exec.Command("cat", "/etc/os-release").Output()
	kernel, _ := exec.Command("uname", "-r").Output()
	uptime, _ := exec.Command("uptime", "-p").Output()

	result["Hostname"] = strings.TrimSpace(string(hostname))
	result["KernelVersion"] = strings.TrimSpace(string(kernel))
	result["OSRelease"] = strings.TrimSpace(string(osRelease))
	result["Uptime"] = strings.TrimSpace(string(uptime))

	return result
}
