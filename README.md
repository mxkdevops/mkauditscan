# mkauditscan
```bash
git init
git add README.md
git commit -m "first commit"
git branch -M main
git remote add origin https://github.com/mxkdevops/mkauditscan.git
git push -u origin main

```
## 🔒 Project: Full Linux Security Audit Scanner (mkauditscan)
Your goal is to build a modular, extensible Linux audit tool that checks the full system for security, configuration, and integrity issues.

## 🎯 Goals of Full Audit Scanner
System Info (hostname, OS, uptime)
User Accounts & SSH (who can log in)
Open Ports & Services
Installed Packages
Firewall/iptables rules
Running Processes
Crontabs & Timed Tasks
Sudoers & Permissions
File Integrity
Rootkit check (import existing)

✅ Step-by-Step Roadmap
🔹 Phase 1: Setup & Project Structure
1. Create folder


mkdir ~/mkauditscan && cd ~/mkauditscan
2. Structure


mkauditscan/
├── main.go
├── modules/
│   ├── info.go
│   ├── users.go
│   ├── ports.go
│   ├── firewall.go
│   ├── crontab.go
│   ├── integrity.go
│   └── rootkit.go      ← (import from previous project)
├── config/
│   └── sample-config.json
├── output/
│   └── audit_report.html
├── go.mod
└── Makefile
🔹 Phase 2: Start with Basic Info Collection

🔹 Phase 3: Add Modules One by One
Each module:

Logs to file or HTML

Can optionally be toggled via config

Adds output to report

Examples:

users.go: get all users, sudo access, last login

ports.go: run ss -tuln or netstat for listening ports

firewall.go: check ufw status or iptables -L

integrity.go: calculate checksums of critical files

crontab.go: check /etc/crontab and crontab -l

### Phase 4: Output as HTML + JSON
Use html/template to output styled reports. Also save .json for machine-readable output.

### Phase 5: Run as Cronjob or Systemd Service
```bash
sudo crontab -e

0 3 * * * /usr/local/bin/mkauditscan
```
Or set up:

```bash
sudo nano /etc/systemd/system/mkauditscan.service
```
### Phase 6: Prepare for CloudWatch
Write all outputs to /var/log/mkauditscan/audit.log

Agent can pick this up in next phase

## ✅ Sample Checklist (Dev Phase)
Task	Status
Create project structure	✅
System info module	✅
User & sudo audit	⏳
Port and service scan	⏳
Cronjobs check	⏳
File integrity module	⏳
Rootkit module reuse	⏳
HTML report builder	⏳
JSON report builder	⏳
Config reader support	⏳
Logging to /var/log	⏳
Systemd or cron support	⏳
CloudWatch-ready	🔜



## ✅ Project Structure

mkauditscan/
├── main.go                 # Entry point: generates and saves audit report
├── modules/
│   └── info.go             # System info collection logic
├── output/                # Where audit reports are saved
├── bin/                   # Compiled binary goes here
├── Makefile               # Build, run, and clean tasks
├── go.mod                 # Go module definition
📦 What’s Included
System Info Collection: hostname, uname -a, uptime

Modular Code: Separated into modules/info.go

Makefile: Simplifies build & run commands

Output Handling: Saves report to output/audit_report.txt

## 🔜 Next Steps
We’ll build features incrementally:

 Add logging and output formatting

 Add file integrity check (MD5/sha256)

 Add cron or systemd setup

 Add config file support

 Add alerting (SMTP, webhook)

 Build README, config sample, Dockerfile