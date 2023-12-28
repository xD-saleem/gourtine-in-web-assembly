build:
	@echo "Building..."
	@cd cmd/wasm && GOOS=js GOARCH=wasm go build -o  ../../assets/json.wasm

start:
	@echo "Starting..."
	@go run cmd/server/main.go
