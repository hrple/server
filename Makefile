.PHONY: code-check setup-build-env test only-test only-code-check

code-check: setup-build-env
	./scripts/static-code-analysis.sh

only-code-check:
	./scripts/static-code-analysis.sh

only-test:
	rm -f bin/code-coverage-report.html
	go test -v -race -coverprofile=bin/code-coverage-report-codecov.out -covermode=atomic ./server && go tool cover -html=bin/code-coverage-report-codecov.out -o bin/code-coverage-report.html

test: code-check
	rm -f bin/code-coverage-report.html
	go test -v -race -coverprofile=bin/code-coverage-report-codecov.out -covermode=atomic ./server && go tool cover -html=bin/code-coverage-report-codecov.out -o bin/code-coverage-report.html

test-coverage-sonarcloud: code-check
	go test -short -coverprofile=bin/code-coverage-report-sonarcloud.out `go list ./... | grep -v vendor/` | grep -v _test`
	go tool cover -func=bin/code-coverage-report-sonarcloud.out

setup-build-env:
	./scripts/setup-build-env.sh

init:
	./scripts/init.sh