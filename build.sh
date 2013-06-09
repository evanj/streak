#!/bin/sh

set -e

go build
echo "Running tests:"
go test
go fmt
go fmt cmd/streak_example.go
go build cmd/streak_example.go
echo "SUCCESS"
