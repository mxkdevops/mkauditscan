// modules/info.go
package modules

import (
    "fmt"
    "os/exec"
    "strings"
)

func CollectSystemInfo() string {
    hostname, _ := exec.Command("hostname").Output()
    osInfo, _ := exec.Command("uname", "-a").Output()
    uptime, _ := exec.Command("uptime", "-p").Output()

    return fmt.Sprintf("System Info:\nHostname: %sOS: %sUptime: %s\n",
        strings.TrimSpace(string(hostname)),
        strings.TrimSpace(string(osInfo)),
        strings.TrimSpace(string(uptime)))
}