.SILENT: ;               # no need for @
.ONESHELL: ;             # recipes execute in same shell
.NOTPARALLEL: ;          # wait for target to finish
.EXPORT_ALL_VARIABLES: ; # send all vars to shell

BINARY_NAME=libmatch

export GO111MODULE=on

all: lint test build

build:
	echo '# $@'
	rm -f $(BINARY_NAME)
	scripts/build.sh $(BINARY_NAME)
	echo ''

lint:
	echo '# $@'
	scripts/lint.sh
	echo ''

lintfix:
	echo '# $@'
	gofmt -l -w -s .
	echo ''

test:
	echo '# $@'
	scripts/test.sh
	echo ''
