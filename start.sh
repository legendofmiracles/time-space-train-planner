#!/usr/bin/env bash
set -a

source deployments/conf.env
PORT=3000 go run ./cmd/main.go
