#!/usr/bin/env bash

set -e

RED='\033[00;31m'
GREEN='\033[00;32m'
RESTORE='\033[0m'

if go test ./...; then
  echo -e "${GREEN}OK${RESTORE}"
else
  echo -e "${RED}FAIL${RESTORE}"
  exit 1
fi
