include env.mk

.PHONY: all clean test verify-commit dev check checklist

all: dev verify-commit buildsucc

buildsucc:
	@echo "Build successfully!"

dev: check test

check: fmt lint

FILES := $(wildcard *.go) $(wildcard */*.go) $(wildcard */*/*.go)
fmt:
	@echo "gofmt (simplify)"
	@gofmt -s -l -w $(FILES)

LINT := win\golangci-lint.exe
#$(LINT) := golangci-lint
lint: golangci-lint
	$(LINT) run -v --disable-all --deadline=3m \
	  --enable=misspell \
	  --enable=ineffassign \
	  --enable=typecheck \
	  --enable=varcheck \
	  --enable=unused \
	  --enable=structcheck \
	  --enable=deadcode \
	  --enable=gosimple

clean:
	go clean -i ./...

# Split tests for CI to run `make test` in parallel.
test: checklist
	go test ./...

EXE=.exe
verify-commit:
	go build -o hooks/$@$(EXE) $@/main.go

golangci-lint: $(LINT)

$(GOPATH)/bin/golangci-lint:
	curl -sfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh| sh -s -- -b $(GOPATH)/bin

env:
	@echo "$(OS)"
	@echo "$(CP)"
	@echo "$(FILES)"
