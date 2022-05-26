GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=globlas
BINARY_UNIX=$(BINARY_NAME)_unix

all: test build

.phony: clean
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

$(BINARY_NAME): *.go clean
	$(GOBUILD) -o $(BINARY_NAME) -v

build: $(BINARY_NAME)

test: clean
	$(GOTEST) -v ./...
	shadow -strict $$(go list ./... | grep -v "api$$")
	staticcheck $$(go list ./... | grep -v "api$$")
	golangci-lint run


run: build
		./$(BINARY_NAME)

gen: generate
generate: 
	go generate ./...