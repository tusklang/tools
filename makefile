GOPATH = $(CURDIR)/../../

.PHONY: test
test: go.mod
	go test ./test/all_test.go -run ''

go.mod: