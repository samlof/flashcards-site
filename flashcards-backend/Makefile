all: bin/app
test: unit-test

PLATFORM=local

.PHONY: bin/app
bin/app:
	@docker build . --target bin

.PHONY: unit-test
unit-test:
	@docker build . --target test