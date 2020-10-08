.PHONY: default
default: test

GOPATH = $(CURDIR)/../../../../

.PHONY: all
all:
	cd ../tusk && $(MAKE)
	cd ../oat && $(MAKE)
	cd ../undra && $(MAKE)

.PHONY: test
test: all
	go test ./test/all_test.go -run ''

.PHONY: clean
clean:
	cd ../tusk && $(MAKE) clean
	cd ../oat && $(MAKE) clean
	cd ../undra && $(MAKE) clean

.PHONY: tag
tag:
	cd ../tusk && git tag $(tag)
	cd ../oat && git tag $(tag)
	cd ../undra && git tag $(tag)
	cd ../goat && git tag $(tag)
	git tag $(tag)