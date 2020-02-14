PACKAGES = $(shell go list ./...)

.PHONY: default
default: build

.PHONY: all
all: build

.PHONY: build
build: check
	go build

.PHONY: check
check: vet lint test

.PHONY: vet
vet:
	go vet $(PACKAGES)

.PHONY: lint
lint:
	golint -set_exit_status $(PACKAGES)

.PHONY: test
test:
	go test -cover $(PACKAGES)

.PHONY: coverage
coverage:
	echo "mode: count" > coverage-all.out
	$(foreach pkg,$(PACKAGES),\
		go test -coverprofile=coverage.out -covermode=count $(pkg);\
		tail -n +2 coverage.out >> coverage-all.out;)
	go tool cover -html=coverage-all.out

.PHONY: clean
clean:
	go clean

.PHONY: deps
deps:
	go get -u golang.org/x/lint/golint

