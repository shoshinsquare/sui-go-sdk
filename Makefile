GOCC = go
GOTEST = GODEBUG=cgocheck=0 $(GOCC) test -v ./... -p 2

default: test

## test:                              run unit tests with a 60s timeout
test:
	$(GOTEST) --timeout 60s

lintci:
	@echo "--> Running linter for code"
	@./build/bin/golangci-lint run --config ./.golangci.yml

## lintci-deps:                       (re)installs golangci-lint to build/bin/golangci-lint
lintci-deps:
	rm -f ./build/bin/golangci-lint
	curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b ./build/bin v1.48.0

