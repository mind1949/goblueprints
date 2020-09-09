.PHONY: test

fmt:
	gofmt -w -s .
	go mod tidy