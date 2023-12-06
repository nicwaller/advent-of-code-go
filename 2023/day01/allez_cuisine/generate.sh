#!/bin/bash
printf "package main\n\nconst input=\`%s\n\`" "$(cat input.txt)" > const.go
