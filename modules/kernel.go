package modules

import (
	"os/exec"
	"strings"
)

func CheckKernelHardening() map[string]interface{} {
	result := make(map[string]interface{})

	uname, _ := exec.Command("uname", "-r").Output()
	result["Kernel Version"] = strings.TrimSpace(string(uname))

	// Common sysctl security checks
	sysctlChecks := map[string]string{
		"kernel.randomize_va_space": "2",
		"net.ipv4.ip_forward":       "0",
		"fs.suid_dumpable":          "0",
		"kernel.kptr_restrict":      "1",
	}

	for key, expected := range sysctlChecks {
		out, _ := exec.Command("sysctl", "-n", key).Output()
		actual := strings.TrimSpace(string(out))
		if actual != expected {
			result[key] = "❌ Expected: " + expected + ", Found: " + actual
		} else {
			result[key] = "✅ " + actual
		}
	}

	return result
}
