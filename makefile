GOPATH = $(CURDIR)/../../../../

.PHONY: test
test:
	cd ../tusk && $(MAKE)
	cd ../oat && $(MAKE)
	cd ../undra && $(MAKE)
	go test ./test/all_test.go -run ''