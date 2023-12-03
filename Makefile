default: all

all: build

.PHONY: build
build:
	CGO_ENABLED=0 go build -a -installsuffix cgo -o aoc-2023 .
