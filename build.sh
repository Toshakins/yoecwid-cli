#!/bin/bash

env GOOS=darwin GOARCH=amd64 go build  -o ./bin/yoecwid-cli-macosx yoecwid-cli
env GOOS=linux GOARCH=amd64 go build -o ./bin/yoecwid-cli-linux64 yoecwid-cli
env GOOS=windows GOARCH=amd64 go build -o ./bin/yoecwid-cli-win yoecwid-cli