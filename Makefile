proto:
	docker run --rm -v $(shell pwd):/workspace -w /workspace ghcr.io/wintbiit/tsf4g:latest -C go -O . client.xml
	docker run --rm -v $(shell pwd):/workspace -w /workspace ghcr.io/wintbiit/tsf4g:latest -C go -O . s0.xml
	docker run --rm -v $(shell pwd):/workspace -w /workspace ghcr.io/wintbiit/tsf4g:latest -C go -O . s1.xml
	protoc --go_out=s99 s99.proto
	go mod tidy
	goimports -w .
	gofmt -s -w .