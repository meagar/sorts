.PHONY: html
html:
	GOOS=js GOARCH=wasm go build -ldflags "-w" -o docs/sorts.wasm ./cmd/main.go

