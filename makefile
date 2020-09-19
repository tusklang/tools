GOPATH = $(CURDIR)/../../../../

.PHONY: test
test: go.mod
	go get -u
	cd ../tusk && $(MAKE)
	cd ../oat && $(MAKE)
	cd ../undra && $(MAKE)
	cd ../goat && $(MAKE)
	go test ./test/all_test.go -run ''

go.mod:
	go mod init

.PHONY: clean
clean:
	-go mod tidy

.PHONY: tag
tag:
	cd ../tusk && git tag $(tag)
	cd ../oat && git tag $(tag)
	cd ../undra && git tag $(tag)
	cd ../goat && git tag $(tag)
	git tag $(tag)