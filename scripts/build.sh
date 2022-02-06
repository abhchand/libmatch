#!/usr/bin/env bash

#
# Build and release `libmatch` for a given OS and Architecture
#

set -e

function set_os_and_arch() {
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
}

# Check if binary name was provided
if [[ -z $1 ]]; then
  echo "Usage: build.sh BINARY_NAME" 1>&2
  exit 1
fi

CUR_DIR=$( cd "$(dirname "${BASH_SOURCE[0]}")" ; pwd -P )

LIBMATCH_BINARY=$1
LIBMATCH_DATE=$(date -u +%Y%m%d)
LIBMATCH_VERSION=$(cat $CUR_DIR/../pkg/version/version.go | grep -m 1 "var version" | sed -e 's/.*\"\(.*\)\"$/\1/g')

# Determine OS and Architecture
#
# If running on CI (Github Actions), $GOOS and $GOARCH will be set by
# the `wangyoucao577/go-release-action` Github Action
#
# If they are not set, we attempt to set these values ourselves
if [[ -z "${GOOS}" || -z "${GOARCH}" ]]; then
  set_os_and_arch
else
  LIBMATCH_OS=$(uname -s)
  LIBMATCH_ARCH=$(uname -m)
fi

# Define colors for pretty output
RED='\033[00;31m'
GREEN='\033[00;32m'
RESTORE='\033[0m'

# Build
echo "Building $LIBMATCH_BINARY..."
echo "Date:          $LIBMATCH_DATE"
echo "Version:       $LIBMATCH_VERSION"
echo "OS:            $LIBMATCH_OS"
echo "Architecture:  $LIBMATCH_ARCH"

go build \
  -ldflags "-s -w -X main.version=${LIBMATCH_DATE}-${LIBMATCH_VERSION}-${LIBMATCH_OS}-${LIBMATCH_ARCH}" \
  -o $LIBMATCH_BINARY cmd/libmatch/main.go

# Post-build output
if [ $? -eq 0 ]; then
  echo "Size:          $(du -h $LIBMATCH_BINARY | cut -f 1)"
  echo -e "${GREEN}OK${RESTORE}"
else
  echo -e "${RED}FAIL${RESTORE}"
  exit 1
fi
