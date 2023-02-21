SHELL:=/bin/bash

clean:
	find . -name '*.gen.go' -delete

gen: clean generate
generate:
	go generate ./...

test: 
	go test -v ./...
	shadow -strict $$(go list ./... | grep -v "api$$")
	staticcheck $$(go list ./... | grep -v "api$$")
	golangci-lint run

