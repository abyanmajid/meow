#!/bin/bash
go test -v ./... -cover -coverprofile=coverage.out
go tool cover -html=coverage.out -o coverage.html