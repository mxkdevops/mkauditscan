// ports.go
package modules

import (
	"os/exec"
)

func GetListeningPorts() string {
	output := "\n--- Listening Ports ---\n"
	cmd := exec.Command("ss", "-tuln")
	res, err := cmd.CombinedOutput()
	if err != nil {
		output += "Error: " + err.Error()
	} else {
		output += string(res)
	}
	return output
}
