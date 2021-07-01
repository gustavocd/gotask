## help: print help messages
.PHONY: help
help:
	@echo 'Usage:'
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' | sed -e 's/^/ /'

# ==============================================================================
# Modules support

## deps-reset: reset dependencies
.PHONY: deps-reset
deps-reset:
	git checkout -- go.mod
	go mod tidy

## tidy: add/remove unused modules
.PHONY: tidy
tidy:
	@echo 'Tidying and verifying module dependencies...'
	go mod tidy
	@echo 'Done!'

## deps-upgrade: upgrades dependencies
.PHONY: deps-upgrade
deps-upgrade:
	# go get $(go list -f '{{if not (or .Main .Indirect)}}{{.Path}}{{end}}' -m all)
	go get -u -t -d -v ./...
	go mod tidy

## deps-cleancache: remove object files and cached files
.PHONY: deps-cleancache
deps-cleancache:
	go clean -modcache

## list: list packages or modules
.PHONY: list
list:
	go list -mod=mod all

# ==============================================================================
# Running tests within the local computer

## test: run tests
.PHONY: test
test:
	@echo 'Running test and linting code...'
	go test -race -v -count=1  ./...
	@echo 'Done!'

# ==============================================================================
# Running app within the local computer

## dev: run the cmd/api application
.PHONY: dev
dev:
	@go run cmd/main.go

# ==============================================================================
# Building the binaries
current_time = $(shell date +"%Y-%m-%d:T%H:%M:%S")
git_description = $(shell git describe --always --dirty --tags --long)
linker_flags = '-s -X main.buildTime=${current_time} -X main.version=${git_description}'

## build: build the binary for current OS and linux server
.PHONY: build
build:
	@echo 'Building cmd/main.go...'
	go build -ldflags=${linker_flags} -o=./bin/gotask ./cmd/main.go
	GOOS=linux GOARCH=amd64 go build -ldflags=${linker_flags} -o=./bin/linux_amd64/gotask ./cmd/main.go
