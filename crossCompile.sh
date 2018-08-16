#!/bin/bash
GOOS=linux 
GOARCH=amd64 
CGO_ENABLED=0 
go build -o quotes

# docker build -t foo/quotes .
# docker run --rm -e PLUGIN_FAILED="dd" foo/quotes