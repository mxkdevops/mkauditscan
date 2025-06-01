// integrity.go
package modules

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

var CriticalFiles = []string{"/bin/ls", "/bin/bash", "/usr/bin/ssh"}

func CheckFileIntegrity() string {
	output := "\n--- File Integrity Check (SHA256) ---\n"
	for _, file := range CriticalFiles {
		f, err := os.Open(file)
		if err != nil {
			output += fmt.Sprintf("%s: ERROR opening file\n", file)
			continue
		}
		h := sha256.New()
		_, err = io.Copy(h, f)
		f.Close()
		if err != nil {
			output += fmt.Sprintf("%s: ERROR reading file\n", file)
			continue
		}
		hash := fmt.Sprintf("%x", h.Sum(nil))
		output += fmt.Sprintf("%s: %s\n", file, hash)
	}
	return output
}