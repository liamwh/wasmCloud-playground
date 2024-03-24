#!/usr/bin/env -S just --justfile
# ^ A shebang isn't required, but allows a justfile to be executed
#   like a script, with `./justfile test`, for example.

set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]
set dotenv-load := true
export CARGO_TERM_COLOR := "always"

# Show available commands
default:
    @just --list --justfile {{justfile()}}

# Start wasmCloud and the wash UI
up:
    cd ./docker && docker compose up -d
    wash ui

# Stop wasmCloud stack
down:
    cd ./docker && docker compose down
