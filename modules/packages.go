package modules

import (
	"os/exec"
	"strings"
)

func GetInstalledPackages() map[string]interface{} {
	result := make(map[string]interface{})

	pkgs, err := exec.Command("dpkg-query", "-W", "-f='${binary:Package}\n'").Output()
	if err == nil {
		result["Packages"] = strings.Split(string(pkgs), "\n")
	} else {
		result["Packages"] = "Could not retrieve package list"
	}

	return result
}
