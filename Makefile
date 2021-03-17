help: ## Show this help.
	@fgrep -h "##" $(MAKEFILE_LIST) | fgrep -v fgrep | sed -e 's/\\$$//' | sed -e 's/##//'

all: linux windows ##

windows: ## Build windows amd64 binary
	GOOS=windows GOARCH=amd64 go build -o bin/raid.exe *.go

linux: ## Build linux amd64 binary
	GOOS=linux GOARCH=amd64 go build -o bin/raid *.go
