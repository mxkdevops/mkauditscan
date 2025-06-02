package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"os"
	"path/filepath"
	"strings"

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

type AuditSummary struct {
	TotalIssues     int
	RedFlags        int
	YellowFlags     int
	GreenFlags      int
	Recommendations []string
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

	// Generate summary from full report
	summary := generateSummary(report)

	// Ensure output directory exists
	outputDir := "output"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Fatalf("Failed to create output directory: %v", err)
	}

	// Write JSON report
	writeJSON(report, filepath.Join(outputDir, "audit_report.json"))

	// Write full detailed HTML report
	writeHTML(report, filepath.Join(outputDir, "audit_report.html"))

	// Write summary HTML report
	writeSummaryHTML(summary, filepath.Join(outputDir, "audit_summary.html"))

	fmt.Println("Audit reports saved to output/")
}

// generateSummary scans the audit report data and counts issues by severity.
// Here we simulate scanning strings for keywords "error", "warning", "ok" in values to assign flags.
func generateSummary(report AuditReport) AuditSummary {
	redFlags := 0
	yellowFlags := 0
	greenFlags := 0
	recommendations := []string{}

	// Helper function to scan map[string]interface{} recursively and count flags
	var scanMap func(m map[string]interface{})
	scanMap = func(m map[string]interface{}) {
		for _, v := range m {
			switch val := v.(type) {
			case string:
				lower := strings.ToLower(val)
				if strings.Contains(lower, "error") || strings.Contains(lower, "critical") || strings.Contains(lower, "fail") {
					redFlags++
				} else if strings.Contains(lower, "warning") {
					yellowFlags++
				} else if strings.Contains(lower, "ok") || strings.Contains(lower, "pass") || strings.Contains(lower, "success") {
					greenFlags++
				}
			case map[string]interface{}:
				scanMap(val) // recurse deeper
			case []interface{}:
				for _, item := range val {
					if subm, ok := item.(map[string]interface{}); ok {
						scanMap(subm)
					}
				}
			}
		}
	}

	// Scan all sections
	scanMap(report.Users)
	scanMap(report.Sudoers)
	scanMap(report.Ports)
	scanMap(report.Firewall)
	scanMap(report.Crontab)
	scanMap(report.Integrity)
	scanMap(report.Kernel)
	scanMap(report.SSH)
	scanMap(report.Services)
	scanMap(report.Hardening)
	scanMap(report.Auth)
	scanMap(report.Packages)
	scanMap(report.AuditD)
	scanMap(report.LogCheck)
	scanMap(report.Network)

	// Total issues are red + yellow
	totalIssues := redFlags + yellowFlags

	// Simple recommendation logic - add custom messages if red flags found
	if redFlags > 0 {
		recommendations = append(recommendations, "Critical issues detected - immediate investigation required.")
	}
	if yellowFlags > 0 {
		recommendations = append(recommendations, "Warnings found - review and address as needed.")
	}
	if redFlags == 0 && yellowFlags == 0 {
		recommendations = append(recommendations, "No critical issues detected. System appears healthy.")
	}

	return AuditSummary{
		TotalIssues:     totalIssues,
		RedFlags:        redFlags,
		YellowFlags:     yellowFlags,
		GreenFlags:      greenFlags,
		Recommendations: recommendations,
	}
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
        a.back { display: block; margin-bottom: 20px; }
    </style>
</head>
<body>
	<a class="back" href="audit_summary.html">&larr; Back to Summary</a>
    <h1>System Audit Report</h1>

    <div class="section"><h2>Users</h2><pre>{{ printf "%#v" .Users }}</pre></div>
    <div class="section"><h2>Sudoers</h2><pre>{{ printf "%#v" .Sudoers }}</pre></div>
    <div class="section"><h2>Ports</h2><pre>{{ printf "%#v" .Ports }}</pre></div>
    <div class="section"><h2>Firewall</h2><pre>{{ printf "%#v" .Firewall }}</pre></div>
    <div class="section"><h2>Crontab</h2><pre>{{ printf "%#v" .Crontab }}</pre></div>
    <div class="section"><h2>Integrity</h2><pre>{{ printf "%#v" .Integrity }}</pre></div>
    <div class="section"><h2>Kernel</h2><pre>{{ printf "%#v" .Kernel }}</pre></div>
    <div class="section"><h2>SSH</h2><pre>{{ printf "%#v" .SSH }}</pre></div>
    <div class="section"><h2>Services</h2><pre>{{ printf "%#v" .Services }}</pre></div>
    <div class="section"><h2>Hardening</h2><pre>{{ printf "%#v" .Hardening }}</pre></div>
    <div class="section"><h2>Authentication Logs</h2><pre>{{ printf "%#v" .Auth }}</pre></div>
    <div class="section"><h2>Installed Packages</h2><pre>{{ printf "%#v" .Packages }}</pre></div>
    <div class="section"><h2>AuditD</h2><pre>{{ printf "%#v" .AuditD }}</pre></div>
    <div class="section"><h2>LogCheck</h2><pre>{{ printf "%#v" .LogCheck }}</pre></div>
    <div class="section"><h2>Network Info</h2><pre>{{ printf "%#v" .Network }}</pre></div>

</body>
</html>
`

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

func writeSummaryHTML(summary AuditSummary, filename string) {
	tmpl := `
<!DOCTYPE html>
<html>
<head>
    <meta charset="UTF-8" />
    <title>Audit Summary</title>
    <style>
        body { font-family: Arial, sans-serif; padding: 20px; }
        .flag { font-weight: bold; padding: 5px 10px; border-radius: 5px; color: white; display: inline-block; margin-right: 10px;}
        .red { background: #e74c3c; }
        .yellow { background: #f39c12; }
        .green { background: #27ae60; }
        .recommendation { margin: 10px 0; padding: 10px; background: #f8f8f8; border-left: 5px solid #3498db; }
        .summary-box { margin-bottom: 20px; }
        a { text-decoration: none; color: #2980b9; font-weight: bold; }
    </style>
</head>
<body>
    <h1>Audit Summary</h1>

    <div class="summary-box">
        <p>Total Issues: <strong>{{ .TotalIssues }}</strong></p>
        <p><span class="flag red">Critical: {{ .RedFlags }}</span>
           <span class="flag yellow">Warnings: {{ .YellowFlags }}</span>
           <span class="flag green">OK: {{ .GreenFlags }}</span>
        </p>
    </div>

    <h2>Top Recommendations</h2>
    {{ if .Recommendations }}
        <ul>
        {{ range .Recommendations }}
            <li class="recommendation">{{ . }}</li>
        {{ end }}
        </ul>
    {{ else }}
        <p>No recommendations at this time.</p>
    {{ end }}

    <p><a href="audit_report.html">View Full Detailed Report &rarr;</a></p>

</body>
</html>
`
	t, err := template.New("summary").Parse(tmpl)
	if err != nil {
		log.Fatalf("Summary HTML template parse error: %v", err)
	}

	f, err := os.Create(filename)
	if err != nil {
		log.Fatalf("Failed to create summary HTML file: %v", err)
	}
	defer f.Close()

	if err := t.Execute(f, summary); err != nil {
		log.Fatalf("Failed to execute summary template: %v", err)
	}
}
