export GO111MODULE=on

.PHONY: build, build_linux, build_windows, clean, test, coverage

build:
	go build cmd/php-pds-skeleton/php-pds-skeleton.go

build_linux:
	env GOOS=linux GOARCH=amd64 go build cmd/php-pds-skeleton/php-pds-skeleton.go

build_windows:
	env GOOS=windows GOARCH=amd64 go build cmd/php-pds-skeleton/php-pds-skeleton.go

clean:
	rm -rf php-pds-skeleton php-pds-skeleton.exe

test:
	go test -v -coverprofile cover.out internal/creator/creator.go internal/creator/creator_test.go

coverage:
	go tool cover -html=cover.out
