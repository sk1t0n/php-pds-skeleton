export GO111MODULE=on

.PHONY: build
build:
	go build cmd/php-pds-skeleton/php-pds-skeleton.go

.PHONY: install
install:
	go install github.com/sk1t0n/php-pds-skeleton@latest

.PHONY: clean
clean:
	rm -rf php-pds-skeleton