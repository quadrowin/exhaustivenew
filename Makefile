SHELL := bash

lint:
	golangci-lint run --timeout 5m0s --allow-parallel-runners

test:
	go test ./pkg/analyzer/...


build.plugin:
	go build -buildmode=plugin plugin/exhaustivenew.go