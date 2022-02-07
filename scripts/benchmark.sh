#!/usr/bin/env bash

#
# Run all benchmarks
#

set -e

go test -bench=.
