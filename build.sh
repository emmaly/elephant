#!/bin/bash

DIR="$(dirname "$(readlink -f -- "$0")")"

cd "$DIR/wasm"
GOARCH=wasm GOOS=js go build -o "$DIR/www/elephant.wasm"

cd "$DIR/go"
GOARCH=amd64 GOOS=linux go build -o "$DIR/go.bin"