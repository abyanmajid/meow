#!/bin/bash
go test -v ./internal/... -cover -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html