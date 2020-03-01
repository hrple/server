.PHONY: code-check setup-build-env test

code-check: setup-build-env
	./scripts/static-code-analysis.sh

test: code-check
	go test -v -race -coverprofile=bin/code-coverage-report-codecov.out -covermode=atomic ./... && go tool cover -html=bin/code-coverage-report-codecov.out -o bin/code-coverage-report.html

test-coverage-sonarcloud: code-check
	go test -short -coverprofile=bin/code-coverage-report-sonarcloud.out `go list ./... | grep -v vendor/`
	go tool cover -func=bin/code-coverage-report-sonarcloud.out

setup-build-env:
	./scripts/setup-build-env.sh

init:
	./scripts/init.sh