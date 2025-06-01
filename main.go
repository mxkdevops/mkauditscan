package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"

	"github.com/mxkdevops/mkauditscan/modules"
)

type AuditReport struct {
	Users     string
	Ports     string
	Firewall  string
	Crontab   string
	Integrity string
}

func main() {
	// Run all modules
	users := modules.GetUsersInfo()
	ports := modules.GetListeningPorts()
	firewall := modules.CheckFirewallStatus()
	crontab := modules.GetCrontabEntries()
	integrity := modules.CheckFileIntegrity()

	// Combine into report
	report := AuditReport{
		Users:     users,
		Ports:     ports,
		Firewall:  firewall,
		Crontab:   crontab,
		Integrity: integrity,
	}

	// Ensure output directory exists
	outputDir := "output"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Write JSON report
	writeJSON(report, filepath.Join(outputDir, "audit_report.json"))

	// Write HTML report
	writeHTML(report, filepath.Join(outputDir, "audit_report.html"))

	fmt.Println("Audit reports saved to output/")
}

func writeJSON(report AuditReport, filename string) {
	data, err := json.MarshalIndent(report, "", "  ")
	if err != nil {
		log.Fatalf("JSON marshal error: %v", err)
	}
	if err := os.WriteFile(filename, data, 0644); err != nil {
		log.Fatalf("Failed to write JSON file: %v", err)
	}
}

func writeHTML(report AuditReport, filename string) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8">
    <title>Audit Report</title>
    <style>
        body { font-family: Arial, sans-serif; padding: 20px; }
        h1, h2 { color: #2c3e50; }
        .section { margin-bottom: 30px; }
        pre { background: #f4f4f4; padding: 10px; border-radius: 5px; white-space: pre-wrap; word-wrap: break-word; }
    </style>
</head>
<body>
    <h1>System Audit Report</h1>

    <div class="section">
        <h2>Users</h2>
        <pre>{{ .Users }}</pre>
    </div>

    <div class="section">
        <h2>Ports</h2>
        <pre>{{ .Ports }}</pre>
    </div>

    <div class="section">
        <h2>Firewall</h2>
        <pre>{{ .Firewall }}</pre>
    </div>

    <div class="section">
        <h2>Crontab</h2>
        <pre>{{ .Crontab }}</pre>
    </div>

    <div class="section">
        <h2>Integrity Check</h2>
        <pre>{{ .Integrity }}</pre>
    </div>
</body>
</html>`

	t, err := template.New("report").Parse(tmpl)
	if err != nil {
		log.Fatalf("HTML template parse error: %v", err)
	}
	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create HTML file: %v", err)
	}
	defer f.Close()

	if err := t.Execute(f, report); err != nil {
		log.Fatalf("Failed to execute template: %v", err)
	}
}
