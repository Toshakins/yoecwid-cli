#!/bin/bash

env GOOS=darwin GOARCH=amd64 go build  -o ./bin/yoecwid-cli-mac-amd64 yoecwid-cli
env GOOS=linux GOARCH=amd64 go build -o ./bin/yoecwid-cli-linux-amd64 yoecwid-cli
env GOOS=windows GOARCH=amd64 go build -o ./bin/yoecwid-cli-windows-amd64 yoecwid-cli