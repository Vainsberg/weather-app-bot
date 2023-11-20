GOCMD = go
GORUN = $(GOCMD) run
BINARY_NAME = weather-app-bot
MAIN_FILE = main.go

run:
	$(GORUN) $(MAIN_FILE)