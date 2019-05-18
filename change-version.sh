#!/bin/bash

function print_usage () {
	echo "  Usage: $0 -v v1.0.0"
	echo
	echo "  -v    --set-version version	for changing the version"
	echo "  -h    --help				for displaying this help"
	echo
	exit 0
}

currentVersion=$(grep 'Version = ' constants/strings.go | awk -F'"' '{print $2}')

if [[ $# -eq 0 ]]; then
    print_usage
fi

while [[ $# -gt 0 ]]
do
	key="$1"

	case $key in
		-v|--set-version)
			VERSION="$2"
			shift 2
			;;
		-h|--help)
			print_usage
			shift # past argument
			;;
	esac
done

function updateVersion() {
	sed -i.bak "s/Version = \".*\"/Version = \"${VERSION}\"/g" constants/strings.go
	sed -i.bak "s/BuildDate = \".*\"/BuildDate = \"$(date +"%Y-%m-%d %H:%M:%S %Z")\"/g" constants/strings.go
	rm  constants/strings.go.bak
}

if [ "X"$VERSION != "X" ]; then
	updateVersion
fi
