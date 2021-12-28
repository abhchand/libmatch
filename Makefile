.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

BINARY_NAME=libmatch

export GO111MODULE=on

all: lint test build

build:
	rm -f $(BINARY_NAME)
	scripts/build.sh $(BINARY_NAME)

lint:
	scripts/gofmt.sh

lintfix:
	go fmt ./...

test:
	go test ./...
