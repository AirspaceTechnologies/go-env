.DEFAULT_GOAL := build

ifneq ("$(wildcard .env.makefile)", "")
       include .env.makefile
endif


GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOLANGCLILINT=golangci-lint

build:
	$(GOBUILD) ./...

clean_test:
	$(GOCLEAN) -testcache

.PHONY: test
test: lint
	$(GOTEST) -race -coverprofile=cover.out ./... && go tool cover -html=cover.out -o cover.html

.PHONY: lint
lint:
	# TODO: fix this and probably replace with revive
	#$(GOLANGCLILINT) run

# Auto-fixes (some) errors
.PHONY: lint_fix
lint_fix:
	$(GOLANGCLILINT) run --fix

%_test:
	cd $* && $(GOTEST) -v

%_test_mem:
	cd $* && $(GOTEST) -memprofile=mem.out

%_test_cpu:
	cd $* && $(GOTEST) -cpuprofile=cpu.out
