.PHONY: code-check setup-build-env test

code-check: setup-build-env
	./scripts/static-code-analysis.sh

test: code-check
	go test -v ./...

test-minimal: code-check
	go test ./...

setup-build-env:
	./scripts/setup-build-env.sh

init:
	git config core.hooksPath .githooks