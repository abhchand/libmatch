#!/usr/bin/env bash

set -e

RED='\033[00;31m'
GREEN='\033[00;32m'
RESTORE='\033[0m'

GOFMTOUTPUT=$(gofmt -l .)
echo -e "${RED}$GOFMTOUTPUT${RESTORE}"

# A "blank" output still outputs a single line
if [[ $(echo $GOFMTOUTPUT | wc -l) -gt 1 ]]; then
  echo "One or more unformatted files found."
  exit 1
else
  echo -e "${GREEN}ok${RESTORE}"
fi
