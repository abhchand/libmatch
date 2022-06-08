#!/usr/bin/env bash

# ########################################################################
# Generates static HTML docs via `godoc`
#
# Apparently this isn't built in to `godoc` yet -_-
#   * https://stackoverflow.com/questions/13530120
#   * https://github.com/golang/go/issues/2381
#
# Adapted from
#   https://gist.github.com/MrWormHole/9eebb821e0aa79f8599920fc0aa34b24
#
# Run `./doc_gen.sh`
# ########################################################################

function extract_module_name {
    # Extract module name
    sed -n -E 's/^\s*module\s+([[:graph:]]+)\s*$/\1/p'
}

function normalize_url {
    # Normalize provided URL. Removing double slashes
    echo "$1" | sed -E 's,([^:]/)/+,\1,g'
}

function generate_go_documentation {
    # Go doc
    local URL
    local PID
    local STATUS

    OUTPUT_DIR=$1

    # Setup
    rm -rf "$OUTPUT_DIR"

    # Extract Go module name from a Go module file
    if [[ -z "$GO_MODULE" ]]; then
        local FILE

        FILE="$(go env GOMOD)"

        if [[ -f "$FILE" ]]; then
            GO_MODULE=$(cat "$FILE" | extract_module_name)
        fi
    fi

    # URL path to Go package and module documentation
    URL=$(normalize_url "http://${GO_DOC_HTTP:-localhost:6060}/pkg/$GO_MODULE/")

    # Starting godoc server
    echo "Starting godoc server..."
    godoc -http="${GO_DOC_HTTP:-localhost:6060}" &
    PID=$!

    # Waiting for godoc server
    while ! curl --fail --silent "$URL" 2>&1 >/dev/null; do
        sleep 0.1
    done

    # Download all documentation content from running godoc server
    wget \
        --recursive \
        --no-verbose \
        --convert-links \
        --page-requisites \
        --adjust-extension \
        --execute=robots=off \
        --include-directories="/lib,/pkg/$GO_MODULE,/src/$GO_MODULE" \
        --exclude-directories="*" \
        --directory-prefix="$OUTPUT_DIR" \
        --no-host-directories \
        "$URL"

    # Stop godoc server
    kill -9 "$PID"
    echo "Stopped godoc server"
    echo "Go source code documentation generated under $OUTPUT_DIR"
}

generate_go_documentation $1
