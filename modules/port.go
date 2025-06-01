package modules

import (
	"os/exec"
)

func GetListeningPorts() map[string]interface{} {
	result := make(map[string]interface{})

	// netstat (preferred: ss or lsof in production)
	output, err := exec.Command("ss", "-tuln").Output()
	if err != nil {
		result["Error"] = "Unable to retrieve port info"
	} else {
		result["ListeningPorts"] = string(output)
	}

	return result
}
