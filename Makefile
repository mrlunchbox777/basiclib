.PHONY: all build test clean

all: build test

build:
	@echo "make building..."
	./scripts/build.sh

test:
	@echo "make testing..."
	./scripts/test.sh

clean:
	@echo "make cleaning..."
	./scripts/clean.sh
