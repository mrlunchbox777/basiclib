.PHONY: all build test clean

all: build test

build:
	@echo "make building..."
	./scripts/build.sh

test:
	@echo "make testing..."
	./scripts/test.sh

coverage:
	@echo "make coverage..."
	./scripts/coverage.sh

clean:
	@echo "make cleaning..."
	./scripts/clean.sh
