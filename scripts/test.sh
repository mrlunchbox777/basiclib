#! /usr/bin/env bash

# Exit on error
set -e

# Get old directory
OLD_DIR="$(pwd)"
ERROR=0

{
	# Get current directory
	DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" >/dev/null 2>&1 && pwd )"

	# Get package directory
	PACKAGE_DIR="$(dirname "$DIR")/pkg"
	cd "$PACKAGE_DIR"

	# Run tests
	echo "Running tests in $PACKAGE_DIR..."
	go test -v -coverprofile=coverage.out -coverpkg=./... ./...
} || {
	ERROR=$?
}

cd "$OLD_DIR"
exit $ERROR