#!/bin/bash

DIR="$(dirname "$(readlink -f -- "$0")")"

cd "$DIR/wasm"
GOARCH=wasm GOOS=js code
