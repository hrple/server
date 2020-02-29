.PHONY: code-check setup-build-env test

code-check: setup-build-env
	./scripts/static-code-analysis.sh

test: code-check
	go test -v -race -coverprofile=bin/code-coverage-report-codecov.out -covermode=atomic ./...

# test-coverage-sonarcloud: code-check
# 	go test -short -coverprofile=bin/code-coverage-report.out `go list ./... | grep -v vendor/`
# 	go tool cover -func=bin/code-coverage-report.out

setup-build-env:
	./scripts/setup-build-env.sh

init:
	git config core.hooksPath .githooks