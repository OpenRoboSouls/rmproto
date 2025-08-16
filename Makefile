proto:
	docker run --rm -v $(shell pwd):/workspace -w /workspace ghcr.io/wintbiit/tsf4g:latest -C go -O . core.xml client.xml
	go mod tidy
	goimports -w .
	gofmt -s -w .