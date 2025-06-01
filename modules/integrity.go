package modules

import (
	"os/exec"
	"strings"
)

func CheckFileIntegrity() map[string]interface{} {
	result := make(map[string]interface{})

	// Check for changes in important system files using md5sum (example)
	files := []string{"/etc/passwd", "/etc/shadow", "/etc/ssh/sshd_config"}
	hashes := make(map[string]string)

	for _, file := range files {
		output, err := exec.Command("md5sum", file).Output()
		if err == nil {
			parts := strings.Fields(string(output))
			if len(parts) >= 1 {
				hashes[file] = parts[0]
			}
		} else {
			hashes[file] = "Error hashing"
		}
	}
	result["MD5Hashes"] = hashes

	// Add future tripwire/aide integration here
	return result
}
