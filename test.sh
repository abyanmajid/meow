#!/bin/bash
go test -v ./pkg/... -cover -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html