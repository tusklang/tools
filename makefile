GOPATH = $(CURDIR)/../../../../

.PHONY: test
test:
	cd ../tusk && $(MAKE)
	cd ../oat && $(MAKE)
	cd ../undra && $(MAKE)
	go test ./test/all_test.go -run ''

.PHONY: tag
tag:
	cd ../tusk && git tag $(tag)
	cd ../oat && git tag $(tag)
	cd ../undra && git tag $(tag)
	cd ../goat && git tag $(tag)
	git tag $(tag)