proto:
	docker run --rm -v $(shell pwd):/workspace -w /workspace ghcr.io/wintbiit/tsf4g:latest -C go -O . client_proto.xml
	goimports -w client
	gofmt -s -w client