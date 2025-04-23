#!/bin/bash
GOOS=windows GOARCH=amd64 garble -tiny -seed=random -literals build -ldflags -H=windowsgui
