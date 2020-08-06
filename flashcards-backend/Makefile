all: bin/app
test: unit-test
lint: lint

PLATFORM=local

.PHONY: bin/app
bin/app:
	@docker build . --target bin --output bin/

.PHONY: unit-test
unit-test:
	@docker build . --target test

.PHONY: lint
lint:
	@docker build . --target lint