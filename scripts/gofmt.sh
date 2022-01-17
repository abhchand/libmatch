#!/usr/bin/env bash

set -e

RED='\033[00;31m'
GREEN='\033[00;32m'
RESTORE='\033[0m'

if [ "$(gofmt -s -l . | tee /dev/stderr | wc -l)" -gt 0 ]; then
  echo -e "${RED}FAIL${RESTORE}"
  exit 1
else
  echo -e "${GREEN}OK${RESTORE}"
fi
