#!/usr/bin/env bash

set -e

if [[ -z $1 ]]; then
  echo "Usage: build.sh BINARY_NAME" 1>&2
  exit 1
fi

LIBMATCH_BINARY=$1
LIBMATCH_DATE=$(date -u +%y%m%d)
LIBMATCH_VERSION='abcde' #$(git describe --always)

RED='\033[00;31m'
GREEN='\033[00;32m'
RESTORE='\033[0m'

if [[ $OS == 'Windows_NT' ]]; then
  LIBMATCH_OS=win32
  if [[ $PROCESSOR_ARCHITEW6432 == 'AMD64' ]]; then
    LIBMATCH_ARCH=amd64
  else
    if [[ $PROCESSOR_ARCHITECTURE == 'AMD64' ]]; then
      LIBMATCH_ARCH=amd64
    fi
    if [[ $PROCESSOR_ARCHITECTURE == 'x86' ]]; then
      LIBMATCH_ARCH=ia32
    fi
  fi
else
  LIBMATCH_OS=$(uname -s)
  LIBMATCH_ARCH=$(uname -m)
fi

echo "Building $LIBMATCH_BINARY..."
go build \
  -ldflags "-s -w -X main.version=${LIBMATCH_DATE}-${LIBMATCH_VERSION}-${LIBMATCH_OS}-${LIBMATCH_ARCH}" \
  -o $LIBMATCH_BINARY cmd/libmatch/main.go

if [ $? -eq 0 ]; then
  du -h $LIBMATCH_BINARY
  echo -e "${GREEN}OK${RESTORE}"
else
  echo -e "${RED}FAIL${RESTORE}"
  exit 1
fi
