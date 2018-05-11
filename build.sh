#!/bin/bash

platforms=(
    darwin-amd64 dragonfly-amd64 freebsd-amd64 linux-amd64 netbsd-amd64 openbsd-amd64 solaris-amd64 windows-amd64
    freebsd-386 linux-386 netbsd-386 openbsd-386 windows-386
    linux-arm linux-arm64 linux-ppc64 linux-ppc64le
)

minimal=(
    darwin-amd64 linux-amd64 windows-amd64
    linux-386 windows-386
    linux-arm linux-arm64
)

build() {
	go run build.go "$@"
}

go run build.go clean

for plat in "${minimal[@]}"; do
    echo Building "$plat"

    goos="${plat%-*}"
    goarch="${plat#*-}"
    dist="tar"

    if [[ $goos == "windows" ]]; then
        dist="zip"
    fi

    build -goos "$goos" -goarch "$goarch" "$dist"
    echo
done
