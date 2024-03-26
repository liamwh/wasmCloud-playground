#!/usr/bin/env -S just --justfile
# ^ A shebang isn't required, but allows a justfile to be executed
#   like a script, with `./justfile test`, for example.

set windows-shell := ["powershell.exe", "-NoLogo", "-Command"]
set dotenv-load := true

# Show available commands
default:
    @just --list --justfile {{justfile()}}

# Start wasmCloud
up-no-openfga:
    cd ./docker && docker compose up -d

# Start wasmCloud with OpenFGA
up:
    cd ./docker && docker compose -f compose.yaml -f compose.openfga.yaml up -d

# Restart the wasmCloud stack with OpenFGA
restart: down up

# Restart the wasmCloud stack
restart-no-openfga: down up-no-openfga

# Stop wasmCloud stack
down:
    cd ./docker && docker compose down

# Run the wash (wasmCloud Shell) UI dashboard:
ui:
    wash ui
