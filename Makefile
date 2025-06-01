APP_NAME=mkauditscan
OUTPUT_DIR=output
CONFIG=config/sample-config.json

build:
	go build -o $(APP_NAME) main.go users.go ports.go firewall.go crontab.go integrity.go

run: build
	./$(APP_NAME) --config $(CONFIG)

clean:
	rm -f $(APP_NAME)
	rm -f $(OUTPUT_DIR)/audit_report.html
	rm -f $(OUTPUT_DIR)/audit_report.json
