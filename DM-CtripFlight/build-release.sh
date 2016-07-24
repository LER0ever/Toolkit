#!/bin/bash

OSS=(linux)
ARCHS=(amd64)

for os in "${OSS[@]}"; do
    for arch in "${ARCHS[@]}"; do
    	echo "Building for $os($arch)"
        GOOS=$os GOARCH=$arch go build -ldflags "-s -w"
    done
done

OSS=(windows)
ARCHS=(386)

for os in "${OSS[@]}"; do
    for arch in "${ARCHS[@]}"; do
    	echo "Building for $os($arch)"
        GOOS=$os GOARCH=$arch go build -ldflags "-s -w"
    done
done
