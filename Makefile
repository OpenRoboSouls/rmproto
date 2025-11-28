lint:
	buf format -w .
	
proto-client:
	protoc --proto_path=. --go_out=paths=source_relative:. **/*.proto

proto-serialport:
	docker run --rm -v .:/workspace -w /workspace ghcr.io/wintbiit/tsf4g:latest -C go -O . serialport/serialport.xml
	go mod tidy
	goimports -w .
	gofmt -s -w .
