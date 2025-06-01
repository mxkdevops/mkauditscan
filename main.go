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
	Users     map[string]interface{}
	Ports     map[string]interface{}
	Firewall  map[string]interface{}
	Crontab   map[string]interface{}
	Integrity map[string]interface{}
	Sudoers   map[string]interface{}
	Kernel    map[string]interface{}
	SSH       map[string]interface{}
	Hardening map[string]interface{}
	Services  map[string]interface{}
	Auth      map[string]interface{}
	Packages  map[string]interface{}
	AuditD    map[string]interface{}
	LogCheck  map[string]interface{}
	Network   map[string]interface{}
}

func main() {
	// Run all modules
	users := modules.GetUsersInfo()
	sudoers := modules.GetSudoersInfo()
	ports := modules.GetListeningPorts()
	firewall := modules.CheckFirewallStatus()
	crontab := modules.GetCrontabEntries()
	integrity := modules.CheckFileIntegrity()
	kernel := modules.CheckKernelHardening()
	ssh := modules.AuditSSHConfig()
	services := modules.CheckRunningServices()
	hardening := modules.CheckHardening()
	auth := modules.GetAuthInfo()
	packages := modules.GetInstalledPackages()
	auditd := modules.GetAuditdStatus()
	logcheck := modules.GetLogCheckResults()
	network := modules.GetNetworkInfo()

	// Combine into report
	report := AuditReport{
		Users:     users,
		Sudoers:   sudoers,
		Ports:     ports,
		Firewall:  firewall,
		Crontab:   crontab,
		Integrity: integrity,
		Kernel:    kernel,
		SSH:       ssh,
		Services:  services,
		Hardening: hardening,
		Auth:      auth,
		Packages:  packages,
		AuditD:    auditd,
		LogCheck:  logcheck,
		Network:   network,
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

    <div class="section"><h2>Sudoers</h2><pre>{{ .Sudoers }}</pre></div>
<div class="section"><h2>Kernel</h2><pre>{{ .Kernel }}</pre></div>
<div class="section"><h2>SSH</h2><pre>{{ .SSH }}</pre></div>
<div class="section"><h2>Hardening</h2><pre>{{ .Hardening }}</pre></div>
<div class="section"><h2>Services</h2><pre>{{ .Services }}</pre></div>
<div class="section"><h2>Authentication Logs</h2><pre>{{ .Auth }}</pre></div>
<div class="section"><h2>Installed Packages</h2><pre>{{ .Packages }}</pre></div>
<div class="section"><h2>AuditD</h2><pre>{{ .AuditD }}</pre></div>
<div class="section"><h2>LogCheck</h2><pre>{{ .LogCheck }}</pre></div>
<div class="section"><h2>Network Info</h2><pre>{{ .Network }}</pre></div>


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
