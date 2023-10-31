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
	echo "Cleaning $PACKAGE_DIR..."
	go clean
	rm -rf "$PACKAGE_DIR/bin"
} || {
	ERROR=$?
}

cd "$OLD_DIR"
exit $ERROR

