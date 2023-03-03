export GO111MODULE=on

.PHONY: build
build:
	go build cmd/php-pds-skeleton/php-pds-skeleton.go

.PHONY: clean
clean:
	rm -rf php-pds-skeleton

.PHONY: test
test:
	go test -v -coverprofile cover.out internal/creator/creator.go internal/creator/creator_test.go

.PHONY: coverage
coverage:
	go tool cover -html=cover.out
